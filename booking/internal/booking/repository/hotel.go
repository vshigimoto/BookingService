package repository

import (
	"booking/internal/booking/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Repo) CreateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var apartment entity.Hotel
		err := ctx.ShouldBindJSON(&apartment)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		err = r.main.QueryRow("insert into hotel(phone, address,category, rating) values($1, $2, $3, $4) returning ID", apartment.Phone, apartment.Address, apartment.Category, apartment.Rating).Scan(&apartment.Id) // $1 and $2 is prepared statement
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": apartment.Id})
	}
}

func (r *Repo) GetHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apartments := make([]entity.Hotel, 0)

		rows, err := r.replica.Query("SELECT * from hotel")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		for rows.Next() {
			var hotel entity.Hotel
			if err = rows.Scan(&hotel.Id, &hotel.Phone, &hotel.Address, &hotel.Category, &hotel.Rating); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			apartments = append(apartments, hotel)
		}
		if err = rows.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, apartments)
	}
}

func (r *Repo) UpdateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		var apartment entity.Hotel
		// Unmarshal json to a new user
		if err := ctx.ShouldBindJSON(&apartment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		_, err := r.main.Exec("UPDATE hotel SET phone=$1, address=$2, category=$3, rating=$4 WHERE id=$5", apartment.Phone, apartment.Address, apartment.Category, apartment.Rating, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) DeleteHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		_, err := r.main.Exec("DELETE from hotel WHERE id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) GetHotelById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("id")
		if login == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		rows, err := r.replica.Query("SELECT * FROM hotel WHERE id=$1", login)
		var hotel entity.Hotel
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		rows.Next()
		if err = rows.Scan(&hotel.Id, &hotel.Phone, &hotel.Address, &hotel.Category, &hotel.Rating); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, hotel)
	}
}
