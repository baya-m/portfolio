package api

import (
	"log"
	"net/http"

	"encoding/json"

	"fmt"

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

	password := login.Password
	fmt.Println(password)

	//ユーザー情報をDBから取得
	result, err := usecase.LoginUsecase{}.Select(login)
	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to login", http.StatusInternalServerError)
		return
	}
	hassdPassword, err := bcrypt.GenerateFromPassword([]byte(result.Password), 10)

	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "failed to login", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hassdPassword), []byte(password))

	if err != nil {
		log.Printf("Error Reason %v", err)
		APIError(w, "Invalid password", http.StatusBadRequest)
		return
	}
	tokenString, err := usecase.CreateToken{}.Create(login)
	if err != nil {
		log.Printf("Failed to Create Token")
		APIError(w, "failed to Create Token", http.StatusInternalServerError)
		return
	}
	APICreated(w, tokenString, http.StatusCreated)
}
