package controllers

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"movie-api/models"
	"movie-api/repository/user"
	"movie-api/utils"
	"net/http"
)

var registeredUsers []models.RegisteredUser

type Controller struct {
}

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.RegisteredUser
		//var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(hash)
		userRepo := userRepository.UserRepository{}
		user = userRepo.Signup(db, user)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, user)
		//json.NewEncoder(w).Encode(user)
	}
}
