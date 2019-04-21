package userRepository

import (
	"database/sql"
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
	//stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id;"
	var userId int
	err := db.QueryRow("insert into users (email, password, userid) values ($1, $2, $3) returning id;", user.Email, user.Password, user.UserId).Scan(&userId);
	logFatal(err)
	user.Password = ""
	return user
}