package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"movie-api/models"
	"net/http"
	"os"
	"strings"
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

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				//errorObject.Message=error.Error()
				RespondWithError(w, http.StatusUnauthorized, models.Error{error.Error(), 401, "UNAUTHARIZED"})
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				RespondWithError(w, http.StatusUnauthorized, models.Error{error.Error(), 401, "UNAUTHARIZED"})
				return
			}

		} else {
			RespondWithError(w, http.StatusUnauthorized, models.Error{"invalid token", 401, "UNAUTHARIZED"})
			return
		}
	})
}
