package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/portfolio/internal/interfaces/api"
	"github.com/portfolio/pkg/config"
)

func StartWebServer() error {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Access-Control-Allow-Origin"})
	router := mux.NewRouter()
	router.HandleFunc("/api/login", api.LoginHandler).Methods("POST")
	router.HandleFunc("/api/signup", api.SignupHandler).Methods("POST")

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
}
