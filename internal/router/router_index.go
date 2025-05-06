package router

import (
	"github.com/johnfarrell/datastar-playground/internal/pages"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := pages.Index(name).Render(w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
