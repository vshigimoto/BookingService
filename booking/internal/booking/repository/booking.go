package repository

import (
	"booking/internal/booking/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Repo) BookRoom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		rows, err := r.replica.Query("SELECT * FROM bookcalendar WHERE hotel_id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		var bookCalendar entity.BookCalendar
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				return
			}
		}(rows)
		rows.Next()
		if err = rows.Scan(&bookCalendar.Id, &bookCalendar.HotelId, &bookCalendar.RoomCount); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if bookCalendar.RoomCount <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "all rooms reserved"})
		}
		_, err = r.main.Exec("UPDATE bookcalendar SET hotel_id=$1, room_count=$2 WHERE id=$3", bookCalendar.HotelId, bookCalendar.RoomCount-1, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "room reserved"})
	}
}
