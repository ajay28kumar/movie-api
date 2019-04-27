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

//var users []models.User

type Controller struct {
}

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var userDetails models.UserDetails
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" || user.Password == "" {
			error.Message = "Email or Password is missing."
			error.Code = "INCOMPLETE_DATA"
			error.Status = 400
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(hash)
		userRepo := userRepository.UserRepository{}
		userDetails = userRepo.Signup(db, user)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, userDetails)
		//json.NewEncoder(w).Encode(user)
	}
}
