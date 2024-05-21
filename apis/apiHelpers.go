package apis

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	// Not much more to do in this function than log the error if something goes wrong

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error in respondWithJSON marshaling %+v: %v", payload, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("error in respondWithJSON writing the response: %v", err)
	}
}
