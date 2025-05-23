package handlers

import (
	"net/http"
	"spelling-bee/utils"
)

func (h *AppHandler) SpellingHandler(w http.ResponseWriter, r *http.Request) {
	ses, _ := h.DB.GetSessionByID(1)

	utils.RenderTemplate(w, r, "spelling.html", map[string]interface{}{
		"CountdownTime":   ses.AnswerTime,
		"DisplayDuration": ses.DisplayTime * 1000,
	})
}
