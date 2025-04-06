package handlers

import (
	"html/template"
	"net/http"
	"spelling-bee/models"

	"github.com/rs/zerolog/log"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/home.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		Title:      "Home",
		ShowNavbar: false,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Error().Err(err).Msg("Error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
