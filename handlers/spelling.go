package handlers

import (
	"html/template"
	"net/http"
	"spelling-bee/models"

	"github.com/rs/zerolog/log"
)

type SpellingPageData struct {
	models.PageData
	Word     string
	Duration int
}

func SpellingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/spelling.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := SpellingPageData{
		PageData: models.PageData{
			Title:      "Spelling",
			ShowNavbar: true,
		},
		Word:     "FOCUS",
		Duration: 60,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Error().Err(err).Msg("Error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
