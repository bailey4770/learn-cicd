package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, logErr error) {
	if logErr != nil {
		log.Println(logErr)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	/*Gosec marks this func as at risk of XSS via Taint Analysis attack. This is a false positive:
	* - 'Content-Type' header is explicitly set as a non-executable (application/json) before writing.
	* - Payload is sourced either from explicitly written strings in our code, our from database query results.
	* - If unsafe data is stored in the database, this problem is better tackled at the input layer, not here.*/
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(dat); err != nil { // #nosec G705
		log.Printf("Error writing data to response: %v", err)
		w.WriteHeader(500)
		return
	}
}
