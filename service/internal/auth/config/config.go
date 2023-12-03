package config

import "time"

type Config struct {
	HttpServer HttpServer `yaml:"HttpServer"`
	Database   Database   `yaml:"Database"`
	Auth       Auth       `yaml:"JwtSecretKey"`
	Transport  Transport  `yaml:"Transport"`
	Kafka      Kafka      `yaml:"Kafka"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
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

type Auth struct {
	PasswordSecretKey string `yaml:"PasswordSecretKey"`
	JwtSecretKey      string `yaml:"JwtSecretKey"`
}

type Transport struct {
	User     UserTransport     `yaml:"user"`
	UserGrpc UserGrpcTransport `yaml:"userGrpc"`
}

type UserTransport struct {
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

type UserGrpcTransport struct {
	Host string `yaml:"host"`
}

type Kafka struct {
	Brokers  []string `yaml:"brokers"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

type Producer struct {
	Topic string `yaml:"topic"`
}

type Consumer struct {
	Topics []string `yaml:"topics"`
}
