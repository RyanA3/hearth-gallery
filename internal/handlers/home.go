package handlers

import (
	"net/http"
	"github.com/RyanA3/hearth-gallery/internal/models"
	"github.com/RyanA3/hearth-gallery/internal/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := &models.PageData{
		Title:   "Welcome Home",
		Content: "This is the home page.",
	}
	render.RenderTemplate(w, "error.tmpl", data)
}
