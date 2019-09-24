package api

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/usecase"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signup model.Signup
	fmt.Println("次がBody")
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&signup)

	if signup.ID == "" {
		APIError(w, "No id param", http.StatusBadRequest)
		return
	}
	if signup.Password == "" {
		APIError(w, "No password param", http.StatusBadRequest)
		return
	}
	usecase.SignupUsecase{}.Create(signup)
}
