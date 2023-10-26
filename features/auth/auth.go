package auth

type User struct {
	ID				uint
	Name			string
	Username		string
	Password		string
	Email			string
	DateBirth		string
	Gender			string
	Role			string
}

type Rating struct {
	ID 				uint
	Like			string
	Komentar		string
	Save			string
	UserID			string
}

type Kategori struct {
	ID				uint
	JenisKategori	string
}

type Naskah struct {
	ID				uint
	Title			string
	Date			string
	UserID			uint
	RatingID		uint
	KategoriID		uint
}