package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ydubinskyi/chirpy/internal/auth"
	"github.com/ydubinskyi/chirpy/internal/database"
)

type PolkaEventName string

const (
	PolkaUserUpgraded PolkaEventName = "user.upgraded"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if apiKey != cfg.polkaApiKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	if params.Event != string(PolkaUserUpgraded) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userUuid, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID", err)
		return
	}

	_, err = cfg.db.UpdateUserSubscription(r.Context(), database.UpdateUserSubscriptionParams{
		ID:          userUuid,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
