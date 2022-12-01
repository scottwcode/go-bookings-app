package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/scottwcode/bookings-app/pkg/config"
	"github.com/scottwcode/bookings-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// muxer from https://github.com/bmizerany/pat
	// which is a Sinatra style pattern muxer for Go's net/http library
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// muxer from https://github.com/go-chi/chi
	// chi is a lightweight, idiomatic and composable router for building
	// Go HTTP services for connecting to various middleware handlers
	mux := chi.NewRouter()

	// Use Recoverer middleware handler
	//    https://pkg.go.dev/github.com/go-chi/chi/middleware#Recoverer
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
