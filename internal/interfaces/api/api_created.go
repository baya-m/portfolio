package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONToken struct {
	Token interface{} `json:"token"`
}
type JSONMessage struct {
	Message interface{} `json:"message"`
}

func APITokenCreated(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json, err := json.Marshal(JSONToken{
		Token: data,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
func APIMessageCreated(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json, err := json.Marshal(JSONMessage{
		Message: data,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
