package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ConfigurateRouter(r *chi.Mux) {
	configMiddlewares(r)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server healt is already good"))
	})

	configFaults(r)
	configAPIRoutes(r)
}

func configFaults(r *chi.Mux) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
}

func configMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
}

func configAPIRoutes(r *chi.Mux) {
	apiRouter := chi.NewRouter()
	apiRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The api healt is already good"))
	})
	r.Mount("/api", apiRouter)
}
