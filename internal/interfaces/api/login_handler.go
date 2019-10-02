package api

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/portfolio/internal/domain/model"
	"github.com/portfolio/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login model.Login
	var err error

	json.NewDecoder(r.Body).Decode(&login)

	if login.LoginID == "" {
		APIError(w, "No login_id param", http.StatusBadRequest)
		return
	}
	if login.Password == "" {
		APIError(w, "No password param", http.StatusBadRequest)
		return
	}

	//ユーザー情報をDBから取得
	result, err := usecase.LoginUsecase{}.Select(login)
	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to login", http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to login", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(login.Password))

	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "Invalid password", http.StatusBadRequest)
		return
	}
	tokenString, err := usecase.CryptUsecase{}.CreateToken(login)
	if err != nil {
		log.Printf("Failed to Create Token")
		APIError(w, "failed to Create Token", http.StatusInternalServerError)
		return
	}
	APITokenCreated(w, tokenString, http.StatusCreated)
}
