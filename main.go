package main

import (
	"net/http"
	"spelling-bee/handlers"
	"spelling-bee/middleware"

	"github.com/rs/zerolog/log"
)

func main() {
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
	mux.Handle("GET /sessions", http.HandlerFunc(handlers.SessionListHandler))
	mux.Handle("POST /sessions", http.HandlerFunc(handlers.CreateSessionHandler))
	mux.Handle("GET /sessions/new", http.HandlerFunc(handlers.NewSessionHandler))
	mux.Handle("GET /spelling", http.HandlerFunc(handlers.SpellingHandler))

	mux.Handle("GET /api/word", http.HandlerFunc(handlers.GetRandomWordAPI))

	loggedMux := middleware.LoggerMiddleware(mux)

	log.Info().Msg("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
