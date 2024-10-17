package router

import (
	"github.com/go-chi/chi"
	"github.com/santa512/monorepo-microservices/bookings/httphandler"
	"github.com/santa512/monorepo-microservices/bookings/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/bookings", func(r chi.Router) {
		r.Get("/", httphandler.GetBookings(s))
		r.Post("/", httphandler.CreateBooking(s))
		r.Delete("/{id}", httphandler.DeleteBooking(s))
	})

	r.Get("/healthy", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
