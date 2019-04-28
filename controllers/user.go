package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" || user.Password == "" {
			utils.RespondWithError(w, http.StatusBadRequest, models.Error{"Email or Password is missing.", 400, "INCOMPLETE_DATA"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, models.Error{})
			return
		}
		user.Password = string(hash)
		userRepo := userRepository.UserRepository{}
		userDetails, err = userRepo.Signup(db, user)
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
				utils.RespondWithError(w, http.StatusForbidden, models.Error{"User Exist.", 401, "USER_EXIST"})
				return
			}
			utils.RespondWithError(w, http.StatusInternalServerError, models.Error{})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, userDetails)
	}
}

func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var userDetails models.UserDetails
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" || user.Password == "" {
			utils.RespondWithError(w, http.StatusBadRequest, models.Error{"Email or Password is missing.", 400, "INCOMPLETE_DATA"})
			return
		}
		password := user.Password
		userRepo := userRepository.UserRepository{}
		user, err := userRepo.Login(db, user)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondWithError(w, http.StatusBadRequest, models.Error{"The user does not exist", 401, "USER_NOT_FOUND"})
				return
			}
			fmt.Println("err : ", err)
			utils.RespondWithError(w, http.StatusInternalServerError, models.Error{})
			return
		}
		hashedPassword := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, models.Error{"Password is wrong", 401, "INVALID_PASSWORD"})
			return
		}
		userDetails.Email = user.Email
		userDetails.UserId = user.UserId
		userDetails.ID = user.ID
		userDetails.Token, err = utils.GenerateToken(userDetails)
		if err != nil {
			fmt.Println("err : ", err)
			utils.RespondWithError(w, http.StatusInternalServerError, models.Error{})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, userDetails)

	}
}
