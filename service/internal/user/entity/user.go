package entity

import "time"

type User struct {
	Id          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Phone       string    `db:"phone"`
	Login       string    `db:"login"`
	Password    string    `db:"password"`
	IsConfirmed bool      `db:"is_confirmed"`
	IsDeleted   bool      `db:"is_deleted"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
