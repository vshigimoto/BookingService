package entity

import "time"

type Photo struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
