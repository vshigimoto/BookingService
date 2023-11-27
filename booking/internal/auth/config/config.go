package config

import "time"

type Config struct {
	Database   Database   `yaml:"Database"`
	HttpServer HttpServer `yaml:"HttpServer"`
	Transport  Transport  `yaml:"Transport"`
	Auth       Auth       `yaml:"Auth"`
	Kafka      Kafka      `yaml:"Kafka"`
}

type Database struct {
	Main    DbNone `yaml:"Main"`
	Replica DbNone `yaml:"Replica"`
}

type DbNone struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	SslMode  string `yaml:"SslMode"`
}

type HttpServer struct {
	Port int `yaml:"Port"`
}

type Transport struct {
	UserTransport UserTransport `yaml:"User" mapstructure:"User"`
	UserGrpc      UserGrpc      `yaml:"UserGrpc" mapstructure:"UserGrpc"`
}

type UserTransport struct {
	Host    string        `yaml:"Host"`
	Timeout time.Duration `yaml:"Timeout"`
}

type UserGrpc struct {
	Host string `yaml:"host"`
}

type Auth struct {
	PasswordSecretKey string `yaml:"passwordSecretKey"`
	JwtSecretKey      string `yaml:"jwtSecretKey"`
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
