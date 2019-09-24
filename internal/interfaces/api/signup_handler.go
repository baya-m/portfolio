package api

import (
	"encoding/json"
	"net/http"

	"fmt"

	"log"

	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/usecase"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signup model.Signup
	fmt.Println("次がBody")
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&signup)

	if signup.LoginID == "" {
		APIError(w, "No login_id param", http.StatusBadRequest)
		return
	}
	if signup.Password == "" {
		APIError(w, "No password param", http.StatusBadRequest)
		return
	}
	err := usecase.SignupUsecase{}.Create(signup)
	if err != nil {
		log.Fatal("Error Reason %v", err)
		APIError(w, "failed to Create ID", http.StatusInternalServerError)
	}
	APICreated(w, fmt.Sprintf("Account Created"), http.StatusAccepted)
}
