package handlers

import (
	"fmt"
	"net/http"
	"spelling-bee/models"
	"spelling-bee/utils"
	"strconv"
)

func (h *AppHandler) SessionListHandler(w http.ResponseWriter, r *http.Request) {
	sessions, _ := h.DB.GetSessions()
	data := map[string]interface{}{
		"Sessions": sessions,
	}

	utils.RenderTemplate(w, r, "session-list.html", data)
}

func (h *AppHandler) NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "session-new.html", map[string]interface{}{})
}

func (h *AppHandler) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	displayTime, _ := strconv.Atoi(r.FormValue("display_time"))
	answerTime, _ := strconv.Atoi(r.FormValue("answer_time"))
	fmt.Println(name)
	fmt.Println(displayTime)
	fmt.Println(answerTime)
	http.Redirect(w, r, fmt.Sprintf("/sessions/detail?id=%d", 2), http.StatusSeeOther)
}

func (h *AppHandler) SessionDetailHandler(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := r.URL.Query().Get("id")
	sessionID, err := strconv.Atoi(sessionIDStr)
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	session, err := h.DB.GetSessionByID(sessionID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rounds, _ := models.GetRoundsBySessionID(sessionIDStr)
	data := map[string]interface{}{
		"Session": session,
		"Rounds":  rounds,
	}
	utils.RenderTemplate(w, r, "session-detail.html", data)
}
