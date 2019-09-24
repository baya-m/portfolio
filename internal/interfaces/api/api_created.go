package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONCreated struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func APICreated(w http.ResponseWriter, createdMessage string, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	jsonCreated, err := json.Marshal(JSONCreated{Message: createdMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonCreated)
}
