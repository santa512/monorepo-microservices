package router

import (
	"github.com/go-chi/chi"

	"github.com/santa512/monorepo-microservices/showtimes/httphandler"
	"github.com/santa512/monorepo-microservices/showtimes/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/showtimes", func(r chi.Router) {
		r.Get("/", httphandler.GetShowTimes(s))
		r.Post("/", httphandler.CreateShowTime(s))
		r.Get("/{date}", httphandler.GetShowTimeByDate(s))
		r.Delete("/{id}", httphandler.DeleteShowTime(s))
	})

	r.Get("/healthy", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
