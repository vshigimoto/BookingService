package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type BookingRepository interface {
	CreateHotel() gin.HandlerFunc
	GetHotel() gin.HandlerFunc
	UpdateHotel() gin.HandlerFunc
	DeleteHotel() gin.HandlerFunc
	GetHotelById() gin.HandlerFunc
}

type Repository interface {
	BookingRepository
}

type Repo struct {
	main    sql.DB
	replica sql.DB
}

func NewRepository(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    *main,
		replica: *replica,
	}
}
