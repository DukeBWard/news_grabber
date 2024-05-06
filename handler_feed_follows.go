package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dukebward/news_grabber/internal/database"
	"github.com/google/uuid"
)

// The apiCfg here in the beginning meaan thats apiConfig structs can user this function
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// these params are generated from sqlc in order to integrate with the query we wanted
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %s", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
