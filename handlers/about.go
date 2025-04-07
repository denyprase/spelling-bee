package handlers

import (
	"net/http"
	"spelling-bee/utils"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "about.html", map[string]interface{}{})
}
