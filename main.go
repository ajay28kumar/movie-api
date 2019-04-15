package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisteredUser struct {
	ID int `json:"id"`
	UserId string `json:"userId"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type userDetails struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	Age int `json:"age"`
	Status bool `json:isActive`
}

func logFatal (err error){
	if err != nil {
		log.Fatal(err)
	}
}

const (
	host     = "localhost"
	port     = 49904
	user     = "postgres"
	password = "9934238755"
	dbname   = "localhost 1"
)


func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	router := mux.NewRouter();
	router.HandleFunc("/login", doLogin).Methods("POST")
	router.HandleFunc("/register", createUser).Methods("POST")
	router.HandleFunc("/user-details", postUserDetails).Methods("POST")
	log.Fatalln(http.ListenAndServe(":8000", router))
}

func doLogin (w http.ResponseWriter, r *http.Request)  {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Print(user)
}

func createUser (w http.ResponseWriter, r *http.Request)  {
	var newUser RegisteredUser
	out, err := uuid.NewV4()
	logFatal(err)
	json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser)
	fmt.Println(out)
}

func postUserDetails (w http.ResponseWriter, r *http.Request)  {
	var userDetails userDetails
	json.NewDecoder(r.Body).Decode(&userDetails)
	fmt.Print(userDetails)
}
