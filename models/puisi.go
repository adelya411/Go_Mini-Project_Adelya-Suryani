package models

type Puisi struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Judul    string `json:"judul"`
	Tanggal  string `json:"tanggal"`
	Isi      string `json:"isi"`
	User     []User `gorm:"foreignkey:ID"`
	Rating   []Rating `gorm:"foreignkey:ID"`
	Kategori []Kategori `gorm:"foreignkey:ID"`
}
