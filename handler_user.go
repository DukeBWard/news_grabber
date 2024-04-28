package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dukebward/news_grabber/internal/auth"
	"github.com/dukebward/news_grabber/internal/database"
	"github.com/google/uuid"
)

// The apiCfg here in the beginning meaan thats apiConfig structs can user this function
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	// this is enabled because queries is allowed to use it, aka the receiver
	// every http request has a context on it, make sure to use the context on any calls in handler that require context
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user from API: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
