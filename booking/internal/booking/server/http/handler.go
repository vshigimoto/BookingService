package http

import (
	"booking/internal/booking/booking"
)

type EndpointHandler struct {
	bookingService booking.Service
}

func NewEndpointHandler(bookingService booking.Service) *EndpointHandler {
	return &EndpointHandler{
		bookingService: bookingService,
	}
}
