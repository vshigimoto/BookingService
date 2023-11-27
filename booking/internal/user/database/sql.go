package database

import (
	"booking/internal/user/config"
	"database/sql"
	"fmt"
)

func (c Config) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.User,
		c.Password,
		c.DbName,
		c.SslMode,
	)
}

type Config config.DbNone

func New(cfg config.DbNone) (*sql.DB, error) {
	conf := Config(cfg)
	db, err := sql.Open("postgres", conf.dsn())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
