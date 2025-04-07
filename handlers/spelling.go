package handlers

import (
	"net/http"
	"spelling-bee/utils"
)

func SpellingHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "spelling.html", map[string]interface{}{})
}
