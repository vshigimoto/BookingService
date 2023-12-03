package database

import (
	"booking/internal/auth/config"
	"database/sql"
	"fmt"
)

func (c Config) dsn() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.Host,
		c.Port,
		c.DbName,
		c.User,
		c.Password,
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
