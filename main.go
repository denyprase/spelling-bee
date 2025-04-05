package main

import (
	"net/http"
	"spelling-bee/handlers"
	"spelling-bee/middleware"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)

	loggedMux := middleware.LoggerMiddleware(mux)

	log.Info().Msg("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
