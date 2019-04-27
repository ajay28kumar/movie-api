package models

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status""`
	Code    string `json:"code""`
}
