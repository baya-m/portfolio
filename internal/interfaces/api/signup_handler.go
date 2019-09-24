package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/usecase"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var signup model.Signup
	if vars["id"] == "" {
		APIError(w, "No id param", http.StatusBadRequest)
		return
	}
	if vars["password"] == "" {
		APIError(w, "No id password", http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&signup)
	usecase.SignupUsecase{}.Create(signup)
}
