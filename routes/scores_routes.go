package routes

import (
	"database/sql"
	"net/http"

	handlers "../handlers"
	"github.com/go-chi/chi"
)

// Users ...
func Users(db *sql.DB) http.Handler {
	scoresHandler := handlers.NewScoresHandler(db)

	r := chi.NewRouter()

	r.Get("/", scoresHandler.GetAll)
	r.Get("/{id:[0-9]+}", scoresHandler.Get)
	r.Post("/", scoresHandler.Create)
	r.Put("/{id:[0-9]+}", scoresHandler.Update)
	r.Patch("/{id:[0-9]+}", scoresHandler.Update)
	r.Delete("/{id:[0-9]+}", scoresHandler.Delete)

	return r
}
