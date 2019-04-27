package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"movie-api/models"
	"net/http"
	"os"
)

type ErrorDetails struct {
	error models.Error
}

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	var errorData ErrorDetails
	errorData.error = error
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorData)
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
		log.Fatal(err)
	}
	return tokenString, nil
}
