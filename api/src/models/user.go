package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Nombre   string `json:"nombre"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
