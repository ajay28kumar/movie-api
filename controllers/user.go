package controllers

import (
	"database/sql"
	"encoding/json"
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
		userRepo := userRepository.UserRepository{}
		user = userRepo.Signup(db,user)
		//if err != nil {
		//	error.Message = "Server error."
		//	utils.RespondWithError(w, http.StatusInternalServerError, error)
		//	return
		//}
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, user)
		json.NewEncoder(w).Encode(user)
	}
}
