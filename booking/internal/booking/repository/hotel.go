package repository

import (
	"booking/internal/booking/entity"
	"context"
	"database/sql"
	"fmt"
)

func (r *Repo) CreateHotel(ctx context.Context, hotel *entity.Hotel) (id int, err error) {
	err = r.main.QueryRow("insert into hotel(name, phone,address, category, rating) values($1, $2, $3, $4, $5) returning ID", hotel.Name, hotel.Phone, hotel.Address, hotel.Category, hotel.Rating).Scan(&hotel.Id) // $1 and $2 is prepared statement
	if err != nil {
		return 0, fmt.Errorf("cannot query with error: %v", err)
	}
	return hotel.Id, nil
}

func (r *Repo) GetHotels(ctx context.Context) ([]entity.Hotel, error) {
	hotels := make([]entity.Hotel, 0)
	rows, err := r.replica.Query("SELECT * from hotel")
	if err != nil {
		return nil, fmt.Errorf("cannot query with error: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var hotel entity.Hotel
		if err = rows.Scan(&hotel.Id, &hotel.Name, &hotel.Phone, &hotel.Address, &hotel.Category, &hotel.Rating); err != nil {
			return nil, fmt.Errorf("cannot scan query with error: %v", err)
		}
		hotels = append(hotels, hotel)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows with error: %v", err)
	}
	return hotels, nil
}

func (r *Repo) UpdateHotel(ctx context.Context, id string, hotel *entity.Hotel) error {
	_, err := r.main.Exec("UPDATE hotel SET name=$1, phone=$2, address=$3, category=$4, rating=$5 WHERE id=$6", hotel.Name, hotel.Phone, hotel.Address, hotel.Category, hotel.Rating, id)
	if err != nil {
		return fmt.Errorf("cannot query with err:%v", err)
	}
	return nil
}

func (r *Repo) DeleteHotel(ctx context.Context, id string) error {
	_, err := r.main.Exec("DELETE from hotel WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("cannot delete user with err:%v", err)
	}
	return nil
}

func (r *Repo) GetHotelById(ctx context.Context, id string) (*entity.Hotel, error) {
	rows, err := r.replica.Query("SELECT * FROM hotel WHERE id=$1", id)
	var hotel entity.Hotel
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	rows.Next()
	if err = rows.Scan(&hotel.Id, &hotel.Name, &hotel.Phone, &hotel.Address, &hotel.Category, &hotel.Rating); err != nil {
		return nil, fmt.Errorf("cannot scan query with error: %v", err)
	}
	return &hotel, nil
}
