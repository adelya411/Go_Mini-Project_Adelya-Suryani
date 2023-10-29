package models

type Rating struct {
	ID       	uint   `gorm:"primaryKey" json:"id"`
	Like    	string `json:"like"`
	Komentar  	string `json:"komentar"`
	Save      	string `json:"save"`
	Penulis 	[]User
}