package models

type RegisteredUser struct {
	ID       int    `json:"id"`
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}