package entity

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	AccessToken  string
	RefreshToken string
}
