package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kitshy/treasure-backend/api/service"
)

type Routes struct {
	router *chi.Mux
	svc    service.Service
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(r *chi.Mux, svc service.Service) Routes {
	return Routes{
		router: r,
		svc:    svc,
	}
}
