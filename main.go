package main

import (
	"net/http"
	"os"
	"spelling-bee/db"
	"spelling-bee/handlers"
	"spelling-bee/middleware"
	"spelling-bee/models"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal().Err(err).Msg("DSN not found in environment")
	}

	database := db.Init(dsn)
	defer database.Close()

	appDB := &models.DB{Conn: database}
	appHandler := &handlers.AppHandler{DB: appDB}

	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		handlers.HomeHandler(w, r)
	}))

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.Handle("GET /about", http.HandlerFunc(handlers.AboutHandler))
	mux.Handle("GET /sessions", http.HandlerFunc(appHandler.SessionListHandler))
	mux.Handle("POST /sessions", http.HandlerFunc(appHandler.CreateSessionHandler))
	mux.Handle("GET /sessions/new", http.HandlerFunc(appHandler.NewSessionHandler))
	mux.Handle("GET /sessions/detail", http.HandlerFunc(appHandler.SessionDetailHandler))
	mux.Handle("GET /spelling", http.HandlerFunc(appHandler.SpellingHandler))

	mux.Handle("GET /api/word", http.HandlerFunc(handlers.GetRandomWordAPI))

	loggedMux := middleware.LoggerMiddleware(mux)

	log.Info().Msg("Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
