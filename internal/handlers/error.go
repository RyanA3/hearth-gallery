package handlers

import (
	"net/http"
	"github.com/RyanA3/hearth-gallery/internal/models"
	"github.com/RyanA3/hearth-gallery/internal/render"
)

func ErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	data := &models.PageData{
		Title:   "Error 404",
		Content: "The requested page does not exist.",
	}
	render.RenderTemplate(w, "error.tmpl", data)
}
