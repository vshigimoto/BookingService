package entity

type Hotel struct {
	Id       int     `json:"Id"`
	Name     string  `json:"Name"`
	Phone    string  `json:"Phone"`
	Address  string  `json:"Address"`
	Category string  `json:"Category"`
	Rating   float64 `json:"Rating"`
}

type BookCalendar struct {
	Id        int `json:"id"`
	HotelId   int `json:"hotel_id"`
	RoomCount int `json:"room_count"`
}

type BookRequest struct {
	Id      int `json:"Id"`
	HotelId int `json:"hotel_id"`
	Code    int `json:"Code"`
}
