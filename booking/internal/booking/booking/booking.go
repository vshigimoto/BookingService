package booking

import (
	"booking/internal/booking/config"
	"booking/internal/booking/repository"
)

type Service struct {
	rep repository.Repository
	cfg config.Config
}

func NewBookingService(rep repository.Repository, cfg config.Config) *Service {
	return &Service{
		rep: rep,
		cfg: cfg,
	}
}
