package handlers

import (
	"fmt"
	"net/http"
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
	displayTime, errDisplay := strconv.Atoi(r.FormValue("display_time"))
	answerTime, errAnswer := strconv.Atoi(r.FormValue("answer_time"))
	if errDisplay != nil || errAnswer != nil {
		fmt.Println("Error converting display or answer time:", errDisplay, errAnswer)
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}
	session, err := h.DB.InsertSession(name, displayTime, answerTime)
	if err != nil {
		fmt.Println("Error inserting session:", err.Error())
		http.Error(w, "Error inserting session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/sessions/detail?id=%d", session.ID), http.StatusSeeOther)
}

func (h *AppHandler) SessionDetailHandler(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := r.URL.Query().Get("id")
	sessionID, err := strconv.Atoi(sessionIDStr)
	if err != nil {
		fmt.Println("Error converting session ID:", err.Error())
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	session, err := h.DB.GetSessionByID(sessionID)
	if err != nil {
		fmt.Println("Error fetching session:", err.Error())
		http.NotFound(w, r)
		return
	}

	rounds, err := h.DB.GetRoundsBySessionID(sessionID)
	if err != nil {
		fmt.Println("Error fetching rounds:", err.Error())
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"Session": session,
		"Rounds":  rounds,
	}
	utils.RenderTemplate(w, r, "session-detail.html", data)
}
