package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portfolio/api"
	config "github.com/portfolio/pkg/configs"
)

func StartWebServer() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/login", api.LoginHandler).Methods("POST")
	//router.HandleFunc("/api/signup", api)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router)
}
