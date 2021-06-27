package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/purna25/go-webapp/pkg/config"
	"github.com/purna25/go-webapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurve)
	mux.Use(SessionLoad)
	// mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
