package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Helper function to handle JSON responses so we do not need to rewrite this code each time we want to send JSON in the response

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

// same for errors

type errResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}
