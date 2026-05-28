package main

import (
	"fmt"
	"net/http"

	"github.com/Aayushman-nvm/RSS-aggregator/db/auth"
	"github.com/Aayushman-nvm/RSS-aggregator/db/sqlc"
)

type authedHandler func(http.ResponseWriter, *http.Request, sqlc.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth errors: %v", err))
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
