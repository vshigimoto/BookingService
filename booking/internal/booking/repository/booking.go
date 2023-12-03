package repository

import (
	"booking/internal/booking/entity"
	"database/sql"
	"fmt"
)

func (r *Repo) BookRoom(id string) error {
	rows, err := r.replica.Query("SELECT * FROM bookcalendar WHERE hotel_id=$1", id)
	if err != nil {
		return fmt.Errorf("cannot query with err :%v", err)
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
		return fmt.Errorf("cannot scan with err :%v", err)
	}
	if bookCalendar.RoomCount <= 0 {
		return fmt.Errorf("all rooms are busy")
	}
	_, err = r.main.Exec("UPDATE bookcalendar SET hotel_id=$1, room_count=$2 WHERE id=$3", bookCalendar.HotelId, bookCalendar.RoomCount-1, id)
	if err != nil {
		return fmt.Errorf("cannot query with err :%v", err)
	}
	return nil
}
