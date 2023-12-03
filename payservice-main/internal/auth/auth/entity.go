package auth

type GenerateTokenRequest struct {
	Login    string
	Password string
}

type JwtUserToken struct {
	AccessToken  string
	RefreshToken string
}

type Code struct {
	UCode string `json:"code"`
}
