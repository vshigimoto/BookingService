package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
			r.Route("/api/auth/v1", func(r chi.Router) {
				r.Use(eh.metricsHandler)
				r.Post("/register", eh.Register)
				r.Post("/user-confirm", eh.ConfirmUser)
				r.Post("/login", eh.Login)
				r.Post("/renew-token", eh.RenewToken)
			})

			r.Handle("/metrics", promhttp.Handler())
		})

	return router
}
