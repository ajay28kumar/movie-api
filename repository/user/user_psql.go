package userRepository

import (
	"database/sql"
	"fmt"
	"log"
	"movie-api/models"
)

type UserRepository struct {}

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func (c UserRepository) Signup(db *sql.DB, user models.RegisteredUser) models.RegisteredUser {
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id;"
	err := db.QueryRow(stmt, user.Email, user.Password, user.UserId).Scan(&user.ID);
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		}else {
			fmt.Println("Else condition")
			logFatal(err);
		}
	}
	user.Password = ""
	return user
}