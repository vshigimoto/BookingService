package entity

type User struct {
	Id          int    `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Phone       string `db:"phone"`
	Login       string `db:"login"`
	Password    string `db:"password"`
	IsConfirmed bool   `db:"isconfirmed"`
}
