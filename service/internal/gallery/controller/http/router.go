package http

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"service/internal/gallery/controller/http/middleware"
)

type router struct {
	logger         *zap.SugaredLogger
	authMiddleware *middleware.JwtV1
}

func NewRouter(logger *zap.SugaredLogger, authMiddleware *middleware.JwtV1) *router {
	return &router{
		logger:         logger,
		authMiddleware: authMiddleware,
	}
}

func (s *router) GetHandler(eh *EndpointHandler) http.Handler {
	router := chi.NewRouter().
		Group(func(r chi.Router) {
			r.Use(s.authMiddleware.AuthV1)
			r.Route("/api/gallery/v1", func(r chi.Router) {
				r.Get("/photo", eh.GetPhotos)
			})
		})

	return router
}

func uploadMultipleFiles(w http.ResponseWriter, r *http.Request) {
	// key = "file" in the form
	for _, files := range r.MultipartForm.File {
		for _, file := range files {
			dst, _ := os.Create(filepath.Join("./", file.Filename))
			f, _ := file.Open()
			io.Copy(dst, f)
		}
	}
}
