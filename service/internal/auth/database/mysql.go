package database

import (
	"context"
	"database/sql"
	"fmt"

	"service/internal/auth/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	db *sqlx.DB
}

type Config config.DbNode

func (c Config) dsn() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func New(cfg config.DbNode) (*Db, error) {
	conf := Config(cfg)

	dbConn, err := sqlx.Connect("mysql", conf.dsn())
	if err != nil {
		return nil, fmt.Errorf("db connect failed: %w", err)
	}

	return &Db{
		db: dbConn,
	}, nil
}

func (d *Db) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	txx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	tCtx := withTx(ctx, txx)
	defer func() {
		expired := fromContext(tCtx)
		if expired != nil {
			expired.IsActive = false
		}
	}()

	if err := fn(tCtx); err != nil {
		terr := txx.Rollback()
		if terr != nil {
			return fmt.Errorf("%w: rollback error: %v", err, terr)
		}
		return err
	}

	return txx.Commit()
}

func (d *Db) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.GetContext(ctx, dest, query, args...)
	}
	return d.db.GetContext(ctx, dest, query, args...)
}

func (d *Db) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.ExecContext(ctx, query, args...)
	}

	return d.db.ExecContext(ctx, query, args...)
}

func (d *Db) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.QueryContext(ctx, query, args...)
	}
	return d.db.QueryContext(ctx, query, args...)
}

func (d *Db) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.SelectContext(ctx, dest, query, args...)
	}
	return d.db.SelectContext(ctx, dest, query, args...)
}

func (d *Db) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.QueryRowxContext(ctx, query, args...)
	}
	return d.db.QueryRowxContext(ctx, query, args...)
}

func (d *Db) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.QueryxContext(ctx, query, args...)
	}
	return d.db.QueryxContext(ctx, query, args...)
}

func (d *Db) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	tx := fromContext(ctx)
	if tx != nil && tx.IsActive {
		return tx.Tx.NamedExecContext(ctx, query, arg)
	}
	return d.db.NamedExecContext(ctx, query, arg)
}

func (d *Db) PingContext(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

func (d *Db) Ping() error {
	return d.db.Ping()
}

func (d *Db) Close() error {
	return d.db.Close()
}

var ctxKey = &transactionCtxKey{}

type TX struct {
	Tx       *sqlx.Tx
	IsActive bool
}

type transactionCtxKey struct{}

func fromContext(ctx context.Context) *TX {
	v := ctx.Value(ctxKey)
	if v == nil {
		return nil
	}
	return v.(*TX)
}

func withTx(parent context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(parent, ctxKey, &TX{Tx: tx, IsActive: true})
}
