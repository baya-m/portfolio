package api

import (
	"encoding/json"
	"net/http"
)

type JSONCreated struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func APICreated(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
