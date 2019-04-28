package userRepository

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
	"movie-api/models"
	"movie-api/utils"
)

type UserRepository struct{}

func (c UserRepository) Signup(db *sql.DB, user models.User) (models.UserDetails, error) {
	u4, err := uuid.NewV4()
	var userDetails models.UserDetails
	if err != nil {
		return userDetails, err
	}
	userId := u4.String()
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id,userid,email"
	row := db.QueryRow(stmt, user.Email, user.Password, userId)
	err = row.Scan(&userDetails.ID, &userDetails.UserId, &userDetails.Email)
	if err != nil {
		return userDetails, err
	}
	userDetails.Token, err = utils.GenerateToken(userDetails)
	if err != nil {
		return userDetails, err
	}
	return userDetails, nil
}

func (c UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.UserId, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
