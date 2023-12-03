package entity

import "time"

type SecretCode struct {
	Id        int       `db:"id"`
	Code      string    `db:"code"`
	UserId    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
