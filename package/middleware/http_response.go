package middleware

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	WriteJSONResponse(w, data, statusCode)
}

func WriteJSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if response != nil {
		json.NewEncoder(w).Encode(response)
	}

}
