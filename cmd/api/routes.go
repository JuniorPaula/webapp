package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewMux()

	// register middleware
	mux.Use(middleware.Recoverer)

	// mux.Use(app.enableCORS)

	// authetication routes - auth handler, refresh
	mux.Post("/auth", app.autheticate)
	mux.Post("/refresh-token", app.refresh)

	// test handler

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// user auth middleware

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/{userID}", app.insertUser)
		mux.Patch("/{userID}", app.updateUser)
	})

	return mux
}
