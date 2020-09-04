package routes

import (
	"database/sql"
	"net/http"

	handlers "../handlers"
	"github.com/go-chi/chi"
)

// Users ...
func Users(db *sql.DB) http.Handler {
	usersHandler := handlers.NewUserHandler(db)

	r := chi.NewRouter()

	r.Get("/", usersHandler.GetAll)
	r.Get("/{id:[0-9]+}", usersHandler.Get)
	r.Post("/", usersHandler.Create)
	r.Put("/{id:[0-9]+}", usersHandler.Update)
	r.Patch("/{id:[0-9]+}", usersHandler.Update)
	r.Delete("/{id:[0-9]+}", usersHandler.Delete)

	return r
}
