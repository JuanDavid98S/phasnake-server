package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	routes "./routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	cors "github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func routerInstace(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(rt chi.Router) {
		rt.Mount("/scores", routes.Scores(db))
	})

	return r
}

func database() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root@192.168.1.7:26257/phasnake?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
		os.Exit(-1)
	}
	return db
}

func main() {
	db := database()
	r := routerInstace(db)

	http.ListenAndServe(":3333", r)
}
