package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"movie-api/controllers"
	"movie-api/driver"
	"net/http"
)


var db *sql.DB
func init() {
	gotenv.Load()
}
func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}
	//router.HandleFunc("/login", doLogin).Methods("POST")
	router.HandleFunc("/sign-up", controller.Signup(db)).Methods("POST")
	//router.HandleFunc("/user-details", postUserDetails).Methods("POST")
	log.Fatalln(http.ListenAndServe(":8000", router))
}
