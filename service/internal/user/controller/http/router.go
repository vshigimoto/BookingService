package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type router struct {
	logger *zap.SugaredLogger
}

func NewRouter(logger *zap.SugaredLogger) *router {
	return &router{
		logger: logger,
	}
}

func (s *router) GetHandler(eh *EndpointHandler) http.Handler {
	router := chi.NewRouter().
		Group(func(r chi.Router) {
			r.Route("/api/user/v1", func(r chi.Router) {
				r.Get("/user/{login}", eh.GetUser())
			})

			r.Handle("/swagger/*", eh.Swagger())
		})

	return router
}
