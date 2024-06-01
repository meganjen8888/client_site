package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/meganjen8888/acupuncture/internal/config"
	"github.com/meganjen8888/acupuncture/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/services", handlers.Repo.Services)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-appointment", handlers.Repo.MakeAppointment)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/head", handlers.Repo.Head)
	mux.Get("/hand", handlers.Repo.Hand)
	mux.Get("/feet", handlers.Repo.Feet)
	mux.Get("/legs", handlers.Repo.Legs)
	mux.Get("/other", handlers.Repo.Other)
	mux.Get("/torso", handlers.Repo.Torso)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
