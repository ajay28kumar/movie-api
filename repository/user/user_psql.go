package userRepository

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"movie-api/models"
	"movie-api/utils"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c UserRepository) Signup(db *sql.DB, user models.User) models.UserDetails {
	u4, err := uuid.NewV4()
	logFatal(err)
	userId := u4.String()
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id,userid,email"
	row := db.QueryRow(stmt, user.Email, user.Password, userId)
	var userDetails models.UserDetails
	err = row.Scan(&userDetails.ID, &userDetails.UserId, &userDetails.Email)
	logFatal(err)
	userDetails.Token, err = utils.GenerateToken(userDetails)
	logFatal(err)
	return userDetails
}
