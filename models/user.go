package models

type User struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	Password string `json:"password"`
}
