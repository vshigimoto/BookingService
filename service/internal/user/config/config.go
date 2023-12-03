package config

import "time"

type Config struct {
	HttpServer HttpServer `yaml:"HttpServer"`
	GrpcServer GrpcServer `yaml:"GrpcServer"`
	Database   Database   `yaml:"Database"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type GrpcServer struct {
	Port string `yaml:"Port"`
}

type Database struct {
	Main    DbNode `yaml:"Main"`
	Replica DbNode `yaml:"Replica"`
}

type DbNode struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}
