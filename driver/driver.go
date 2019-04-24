package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	//host     = "localhost"
	port     = 5432
	user     = "postgres"
	//password = "9934238755"
	//dbname   = "postgres"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), port,user, os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, err := sql.Open("postgres", psqlInfo)
	logFatal(err)
	//defer db.Close()
	err = db.Ping()
	logFatal(err)
	fmt.Println("Successfully connected to db!")
	return db
}
