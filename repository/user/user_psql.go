package userRepository

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"movie-api/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c UserRepository) Signup(db *sql.DB, user models.RegisteredUser) models.RegisteredUser {
	u4,err := uuid.NewV4()
	logFatal(err)
	userId:= u4.String()
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id,userid"
	row := db.QueryRow(stmt, user.Email, user.Password, userId)
	err = row.Scan(&user.ID, &user.UserId)
	logFatal(err)
	user.Password = ""
	return user
}
