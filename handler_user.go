package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Aayushman-nvm/RSS-aggregator/db/sqlc"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON in handler_user: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), sqlc.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
} // takes apicfg as parameter, turned it into a method and now its attached to apiConfig struct as method
