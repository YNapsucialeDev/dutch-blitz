package models

type Game struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
}
