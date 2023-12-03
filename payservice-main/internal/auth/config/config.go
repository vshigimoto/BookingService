package config

type Config struct {
	Auth      Auth
	Redis     Redis
	Database  Database
	Transport Transport
	Kafka     Kafka
}

type Redis struct {
	Address string
}

type Transport struct {
	UserGrpc UserGrpcTransport
}

type UserGrpcTransport struct {
	Host string
}

type Auth struct {
	JwtSecretKey      string
	PasswordSecretKey string
}

type Database struct {
	Main    Node
	Replica Node
}

type Node struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Kafka struct {
	Brokers  []string
	Producer Producer
	Consumer Consumer
}

type Producer struct {
	Topic string
}

type Consumer struct {
	Topics []string
}
