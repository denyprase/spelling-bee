package handlers

import (
	"html/template"
	"net/http"
	"spelling-bee/models"

	"github.com/rs/zerolog/log"
)

type SessionPageData struct {
	models.PageData
	Sessions []string
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/session.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sesh := []string{"Session 1", "Session 2", "Session 3"}

	pageData := models.PageData{
		Title:      "Home",
		ShowNavbar: true,
	}

	data := SessionPageData{
		PageData: pageData,
		Sessions: sesh,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Error().Err(err).Msg("Error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
