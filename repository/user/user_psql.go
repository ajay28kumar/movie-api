package userRepository

import (
	"database/sql"
	"log"
	"movie-api/helper"
	"movie-api/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c UserRepository) Signup(db *sql.DB, user models.RegisteredUser) models.RegisteredUser {
	userId := helper.GenerateUserID()
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id,userid"
	row := db.QueryRow(stmt, user.Email, user.Password, userId)
	err := row.Scan(&user.ID, &user.UserId)
	logFatal(err)
	user.Password = ""
	return user
}
