package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/utkarsh/bookings/pkg/config"
	"github.com/utkarsh/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(Writetoconsole)
	mux.Use(Nosurf)
	mux.Use(Sessionload)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./Static/"))
	mux.Handle("/Static/*", http.StripPrefix("/Static", fileServer))
	return mux
}
