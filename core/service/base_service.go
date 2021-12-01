package service

import (
	"encoding/json"
	"net/http"

	"github.com/ertugrul-k/goap/models/request"
)

// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }

func respondWithJson(w http.ResponseWriter, code int, message string, payload interface{}) error {
	w.WriteHeader(code)
	httpResponse := request.NewResponse(code, message, payload)
	err := json.NewEncoder(w).Encode(httpResponse)
	return err
}
