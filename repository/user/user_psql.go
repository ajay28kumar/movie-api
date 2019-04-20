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
	stmt := "insert into users (email, password, userid) value ($1, $2, $3) returning id;"
	err := db.QueryRow(stmt, user.Email, user.Password, user.UserId).Scan(&user.ID)
	logFatal(err);
	user.Password = ""
	return user
}