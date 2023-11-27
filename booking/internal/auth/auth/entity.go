package auth

type GenerateTokenRequest struct {
	Login    string
	Password string
}

type JwtUserToken struct {
	Token        string
	RefreshToken string
}

type JwtTokenContent struct {
	UserId int
	Name   string
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
