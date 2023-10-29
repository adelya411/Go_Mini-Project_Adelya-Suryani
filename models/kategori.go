package models

type Kategori struct {
	ID       		uint   	`gorm:"primaryKey" json:"id"`
	Jenis_Kategori	string	`gorm:"not null" json:"jeniskategori"`
}