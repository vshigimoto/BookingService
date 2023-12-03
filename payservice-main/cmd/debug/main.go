package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/debug", middleware.Profiler())

	err := http.ListenAndServe(":8083", r)
	if err != nil {
		return
	}
}
