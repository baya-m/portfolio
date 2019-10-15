package api

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/usecase"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signup model.Signup
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "POST")

	json.NewDecoder(r.Body).Decode(&signup)

	if 

	if signup.LoginID == "" {
		APIError(w, "No login_id param", http.StatusBadRequest)
		return
	}
	if signup.Password == "" {
		APIError(w, "No password param", http.StatusBadRequest)
		return
	}
	hash, err := usecase.CryptUsecase{}.HashPassword(signup)
	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to Create ID", http.StatusInternalServerError)
		return
	}

	signup.Password = string(hash)
	err = usecase.SignupUsecase{}.Create(signup)

	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to Create ID", http.StatusInternalServerError)
		return
	}

	APIMessageCreated(w, "Account Created", http.StatusCreated)
	log.Printf("Account Created LoginID : %v", signup.LoginID)
}
