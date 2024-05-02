package main

import (
	"fmt"
	"net/http"

	"github.com/dukebward/news_grabber/internal/auth"
	"github.com/dukebward/news_grabber/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) midAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, user)
	}
}
