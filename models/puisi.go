package models

type Puisi struct {
	ID       	uint   `gorm:"primaryKey" json:"id"`
	Judul    	string `json:"judul"`
	Tanggal  	string `json:"tanggal"`
	Isi      	string `json:"isi"`
	Penulis 	[]User
	Rating		[]Rating
	Kategori	[]Kategori
}