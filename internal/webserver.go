package internal

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portfolio/api"
	config "github.com/portfolio/pkg/configs"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func StartWebServer() error {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/login", api.LoginHandler).Methods("POST")
	fmt.Printf("Hello")
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router)
}
