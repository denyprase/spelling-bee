package utils

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Determine whether navbar should be shown based on the current path
	hiddenPaths := map[string]bool{
		"/": true,
	}
	showNavbar := !hiddenPaths[r.URL.Path]

	// Wrap the data
	templateData := map[string]interface{}{
		"ShowNavbar":  showNavbar,
		"CurrentPath": r.URL.Path,
	}

	// Add user-provided data
	if userData, ok := data.(map[string]interface{}); ok {
		for k, v := range userData {
			templateData[k] = v
		}
	} else if data != nil {
		templateData["Data"] = data
	}

	// Parse only layout and the specific page template
	files := []string{
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", name),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout.html", templateData)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}
