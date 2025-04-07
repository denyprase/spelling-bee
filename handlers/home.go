package handlers

import (
	"net/http"
	"spelling-bee/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "home.html", map[string]interface{}{})
}
