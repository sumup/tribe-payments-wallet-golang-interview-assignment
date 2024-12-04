package api

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	mux *chi.Mux,
) {
	mux.Get("/live", Health)

	mux.Route("/v1", func(r chi.Router) {
	})
}
