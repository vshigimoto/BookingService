package repository

import (
	"context"
	"database/sql"
	"github.com/vshigimoto/BookingService/internal/booking/entity"
)

type BookingRepository interface {
	CreateHotel(ctx context.Context, hotel *entity.Hotel) (id int, err error)
	GetHotels(ctx context.Context) ([]entity.Hotel, error)
	UpdateHotel(ctx context.Context, id string, hotel *entity.Hotel) error
	DeleteHotel(ctx context.Context, id string) error
	GetHotelById(ctx context.Context, id string) (*entity.Hotel, error)
}

type Repository interface {
	BookingRepository
}

type Repo struct {
	main    *sql.DB
	replica *sql.DB
}

func NewRepository(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    main,
		replica: replica,
	}
}
