package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"payservice/internal/user/config"
)

type Config config.Node

func (c Config) dataSourceName() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
	)
}

//cfg config.Database

func Connect(cfg config.Node) *sql.DB {
	c := Config(cfg)
	//dataSourceName := "host=localhost user=postgres password=12345 dbname=payservice port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("postgres", c.dataSourceName())
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
