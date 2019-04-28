package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"movie-api/models"
	"net/http"
	"os"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	if status == 0 || error.Message == "" || error.Code == "" || error.Status == 0 {
		error = serverError()
		w.WriteHeader(error.Status)
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func serverError() models.Error {
	var err models.Error
	err.Status = 500
	err.Code = "NETWORK_ERROR"
	err.Message = "Networ error."
	return err
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func GenerateToken(user models.UserDetails) (string, error) {
	var err error
	secret := os.Getenv("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
