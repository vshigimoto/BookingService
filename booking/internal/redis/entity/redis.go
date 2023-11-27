package entity

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       string `json:"DB"`
}
