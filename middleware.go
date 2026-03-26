package main

import (
	"fmt"
	"net/http"

	"github.com/zavista/rss-feed/internal/auth"
	"github.com/zavista/rss-feed/internal/database"
)

// any auth handler will have this function signature (ex. getUser)
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	// normal http handler that does some middleware (in this case, does the authentication), then runs the handler afterwards with the authenticated user
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
