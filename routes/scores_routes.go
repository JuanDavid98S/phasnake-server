package routes

import (
	"database/sql"
	"net/http"

	handlers "../handlers"
	"github.com/go-chi/chi"
)

// Scores is a method to define routes to scores
func Scores(db *sql.DB) http.Handler {
	scoresHandler := handlers.NewScoresHandler(db)

	r := chi.NewRouter()

	r.Get("/", scoresHandler.GetAll)
	r.Get("/{id}", scoresHandler.Get)
	r.Post("/", scoresHandler.Create)
	r.Put("/{id}", scoresHandler.Update)
	r.Patch("/{id}", scoresHandler.Update)
	r.Delete("/{id}", scoresHandler.Delete)

	return r
}
