package service

// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }

// func RespondWithJson(w http.ResponseWriter, code int, message string, payload interface{}) error {
// 	w.WriteHeader(code)
// 	httpResponse := request.NewResponse(code, message, payload)
// 	err := json.NewEncoder(w).Encode(httpResponse)
// 	return err
// }
