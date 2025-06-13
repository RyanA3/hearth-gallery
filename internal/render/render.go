package render

import (
	"log"
	"net/http"
	"github.com/RyanA3/hearth-gallery/internal/templates"
)

// RenderTemplate renders an HTML template with the given data.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	ts, ok := templates.TemplateCache[tmpl]
	if !ok {
		log.Printf("The template %s does not exist in the cache", tmpl)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err := ts.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
