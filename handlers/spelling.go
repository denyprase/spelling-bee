package handlers

import (
	"net/http"
	"spelling-bee/models"
	"spelling-bee/utils"
)

func SpellingHandler(w http.ResponseWriter, r *http.Request) {
	ses, _ := models.GetSessionByID("1")

	utils.RenderTemplate(w, r, "spelling.html", map[string]interface{}{
		"CountdownTime":   ses.AnswerTime,
		"DisplayDuration": ses.DisplayTime * 1000,
	})
}
