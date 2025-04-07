package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"spelling-bee/models"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SessionsPageData struct {
	models.PageData
	Sessions []models.Session
}

type SessionPageData struct {
	models.PageData
	Session models.Session
	Rounds  []models.Round
}

func SessionListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/session-list.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sessions, _ := models.GetSessions()

	pageData := models.PageData{
		Title:      "Home",
		ShowNavbar: true,
	}

	data := SessionsPageData{
		PageData: pageData,
		Sessions: sessions,
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

func SessionDetailHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("id")
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	rounds, _ := models.GetRoundsBySessionID(sessionID)

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/session-detail.html")
	if err != nil {
		log.Error().Err(err).Msg("Error parsing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := SessionPageData{
		PageData: models.PageData{
			Title:      "Session Details",
			ShowNavbar: true,
		},
		Session: session,
		Rounds:  rounds,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Error().Err(err).Msg("Error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
