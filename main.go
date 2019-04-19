package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisteredUser struct {
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
	port     = 5432
	user     = "postgres"
	password = "9934238755"
	dbname   = "postgres"
)


func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo);
	db, err := sql.Open("postgres", psqlInfo)
	logFatal(err)
	defer db.Close()
	err = db.Ping()
	logFatal(err)
	fmt.Println("Successfully connected!")
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
