package redis

type Config struct {
	Auth     Auth
	Redis    Redis
	Database Database
}

type Redis struct {
	Address string
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
