package handlers

import (
	"fmt"
	"net/http"
	"spelling-bee/models"
	"spelling-bee/utils"
	"strconv"
)

func SessionListHandler(w http.ResponseWriter, r *http.Request) {
	sessions, _ := models.GetSessions()
	data := map[string]interface{}{
		"Sessions": sessions,
	}

	utils.RenderTemplate(w, r, "session-list.html", data)
}

func NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "session-new.html", map[string]interface{}{})
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	displayTime, _ := strconv.Atoi(r.FormValue("display_time"))
	answerTime, _ := strconv.Atoi(r.FormValue("answer_time"))
	fmt.Println(name)
	fmt.Println(displayTime)
	fmt.Println(answerTime)
	http.Redirect(w, r, fmt.Sprintf("/sessions/detail?id=%d", 2), http.StatusSeeOther)
}

func SessionDetailHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("id")
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	rounds, _ := models.GetRoundsBySessionID(sessionID)
	data := map[string]interface{}{
		"Session": session,
		"Rounds":  rounds,
	}
	utils.RenderTemplate(w, r, "session-detail.html", data)
}
