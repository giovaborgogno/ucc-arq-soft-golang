package models

type User struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}
