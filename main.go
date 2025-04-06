package main

import (
	"net/http"
	"spelling-bee/handlers"
	"spelling-bee/middleware"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		handlers.HomeHandler(w, r)
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/session", handlers.SessionHandler)
	mux.HandleFunc("/spelling", handlers.SpellingHandler)

	mux.HandleFunc("/api/word", handlers.GetRandomWordAPI)

	loggedMux := middleware.LoggerMiddleware(mux)

	log.Info().Msg("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
