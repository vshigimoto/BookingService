package entity

type UserToken struct {
	UserID       int
	AccessToken  string
	RefreshToken string
}

type UserRegister struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Phone       string `json:"Phone"`
	Login       string `json:"Login"`
	Password    string `json:"Password"`
	IsConfirmed bool   `json:"isConfirmed"`
}

type UserCod struct {
	Code string `json:"code"`
}
