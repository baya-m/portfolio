package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["login"] == "" {
		APIError(w, "No id param", http.StatusBadRequest)
		return
	}
	if vars["password"] == "" {
		APIError(w, "No id password", http.StatusBadRequest)
		return
	}
}
