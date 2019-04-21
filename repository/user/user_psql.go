package userRepository

import (
	"database/sql"
	"fmt"
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
	u4, err := uuid.NewV4()
	fmt.Println("u4: ", u4);
	fmt.Printf("type of u4 is : %T\n", u4)
	stmt := "insert into users (email, password, userid) values ($1, $2, $3) returning id;"
	//hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	err = db.QueryRow(stmt, user.Email, user.Password, user.UserId).Scan(&user.ID);
	logFatal(err)
	user.Password = ""
	return user
}
