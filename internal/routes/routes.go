package routes

import (
	"net/http"
	"github.com/RyanA3/hearth-gallery/internal/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Register route handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			handlers.ErrorPageHandler(w, r);
			return
		}

		handlers.HomeHandler(w, r);
	})

	mux.HandleFunc("/home", handlers.HomeHandler)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
