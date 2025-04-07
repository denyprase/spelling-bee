package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"spelling-bee/models"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SessionPageData struct {
	models.PageData
	Sessions []string
}

func SessionListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/session-list.html")
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

func NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/session-new.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pageData := models.PageData{
		Title:      "Home",
		ShowNavbar: true,
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Error().Err(err).Msg("Error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	displayTime, _ := strconv.Atoi(r.FormValue("display_time"))
	answerTime, _ := strconv.Atoi(r.FormValue("answer_time"))
	fmt.Println(name)
	fmt.Println(displayTime)
	fmt.Println(answerTime)

}
