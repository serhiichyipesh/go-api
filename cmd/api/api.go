package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/serhiichyipesh/go-api/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Post("/register", app.registerUserHandler)
		r.Get("/users", app.getAllUsersHandler)
		r.Post("/posts", app.createPostHandler)
		r.Get("/posts", app.getAllPostsHandler)
		r.Get("/posts/{id}", app.getPostByIDHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("listening on %s", app.config.addr)

	return srv.ListenAndServe()
}
