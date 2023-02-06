package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *App) Routes() http.Handler {
	mux := chi.NewRouter()
	app.middleware(mux)
	mux.Get("/", app.Home)
	mux.Post("/login", app.login)
	mux.Get("/refresh", app.refresh)
	mux.Get("/logout", app.logout)
	mux.Route("/admin", func(router chi.Router) {
		app.authMiddleware(mux)
		//mux.Get("/users", app.Users)
		//mux.Get("/user/{id}", app.UserEdit)
		//mux.Post("/user/new", app.UserNew)
		//mux.Patch("/user/{id}", app.UserUpdate)
		//mux.Delete("/user/{id}", app.UserDelete)
	})

	return mux
}

// other routes
