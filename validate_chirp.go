package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const maxChirpLength = 140

func (cfg *apiConfig) handleValidateChirp(w http.ResponseWriter, r *http.Request) {
	type validatePayload struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := validatePayload{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respondWithError(w, 500, "Something went wrong")
		return
	}

	bodyLength := len(params.Body)

	if bodyLength > maxChirpLength {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	type successRes struct {
		Valid bool `json:"valid"`
	}

	resBody := successRes{
		Valid: true,
	}

	respondWithJSON(w, 200, resBody)
}
