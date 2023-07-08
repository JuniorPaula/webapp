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

	mux.Use(app.enableCORS)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	mux.Route("/web", func(r chi.Router) {
		r.Post("/auth", app.autheticate)
		// /refresh-token
		r.Get("/refresh-token", app.refreshUsingCookie)
		// /logout
		r.Get("/logout", app.deleteRefreshCookie)
	})

	// authetication routes - auth handler, refresh
	mux.Post("/auth", app.autheticate)
	mux.Post("/refresh-token", app.refresh)

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// user auth middleware
		mux.Use(app.authRequire)

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
