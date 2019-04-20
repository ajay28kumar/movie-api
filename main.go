package main

import (
	"database/sql"
	"movie-api/controllers"
	"movie-api/driver"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"time"
)

func logFatal(err error) {
	if err != nil {
		fmt.Println("err: ", err)
		log.Fatal(err)
	}
}
var db *sql.DB

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}
	//router.HandleFunc("/login", doLogin).Methods("POST")
	router.HandleFunc("/register", controller.Signup(db)).Methods("POST")
	//router.HandleFunc("/user-details", postUserDetails).Methods("POST")
	log.Fatalln(http.ListenAndServe(":8000", router))
}
