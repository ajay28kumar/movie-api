package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	Email string `json:"email"`
	ID       int    `json:"id"`
	UserId   string `json:"userId"`
	Token string `json:"token"`
}
