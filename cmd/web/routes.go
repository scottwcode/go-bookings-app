package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/scottwcode/bookings-app/pkg/config"
	"github.com/scottwcode/bookings-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// muxer from https://github.com/bmizerany/pat
	// which is a Sinatra style pattern muxer for Go's net/http library
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
