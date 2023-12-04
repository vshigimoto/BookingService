package repository

import (
	"booking/internal/booking/entity"
	"database/sql"
	"fmt"
)

func (r *Repo) BookRoom(id string, code int) (int, error) {
	rows, err := r.replica.Query("SELECT * FROM bookcalendar WHERE hotel_id=$1", id)
	if err != nil {
		return 0, fmt.Errorf("cannot query with err :%v", err)
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
		return 0, fmt.Errorf("cannot scan with err :%v", err)
	}
	if bookCalendar.RoomCount <= 0 {
		return 0, fmt.Errorf("all rooms are busy")
	}
	var bookReq entity.BookRequest
	err = r.main.QueryRow("insert into bookrequest(hotel_id, code) values($1, $2) returning ID", id, code).Scan(&bookReq.Id) // $1 and $2 is prepared statement
	if err != nil {
		return 0, fmt.Errorf("cannot query with error: %v", err)
	}
	return bookReq.Id, nil
}

func (r *Repo) ConfirmBook(request *entity.BookRequest) error {
	var bookRequest entity.BookRequest
	rows, err := r.replica.Query("SELECT * FROM bookrequest WHERE id=$1", request.Id)
	if err != nil {
		return fmt.Errorf("query scan with err :%v", err)
	}
	rows.Next()
	if err = rows.Scan(&bookRequest.Id, &bookRequest.HotelId, &bookRequest.Code); err != nil {
		return fmt.Errorf("cannot scan with err :%v", err)
	}
	rows, err = r.replica.Query("SELECT * FROM bookcalendar WHERE hotel_id=$1", bookRequest.HotelId)
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

	if request.Code != bookRequest.Code {
		return fmt.Errorf("code is not valid")
	}
	_, err = r.main.Exec("UPDATE bookcalendar SET hotel_id=$1, room_count=$2 WHERE id=$3", bookCalendar.HotelId, bookCalendar.RoomCount-1, bookRequest.HotelId)
	if err != nil {
		return fmt.Errorf("cannot query with err :%v", err)
	}
	return nil
}
