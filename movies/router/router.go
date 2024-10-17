package router

import (
	"github.com/go-chi/chi"
	"github.com/santa512/monorepo-microservices/movies/httphandler"
	"github.com/santa512/monorepo-microservices/movies/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/movies", func(r chi.Router) {
		r.Get("/", httphandler.GetMovies(s))
		r.Post("/", httphandler.CreateMovie(s))
		r.Get("/{id}", httphandler.GetMovieById(s))
		r.Delete("/{id}", httphandler.DeleteMovie(s))
	})

	r.Get("/healthy", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
