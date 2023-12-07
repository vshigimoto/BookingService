package http

import (
	"github.com/vshigimoto/BookingService/internal/booking/booking"
)

type EndpointHandler struct {
	bookingService booking.Service
}

func NewEndpointHandler(bookingService booking.Service) *EndpointHandler {
	return &EndpointHandler{
		bookingService: bookingService,
	}
}
