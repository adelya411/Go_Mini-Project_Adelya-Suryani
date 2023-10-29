package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID			uint	`gorm:"primary_key" json:"id"`
	Username 	string 	`gorm:"not null" valid:"required~your username is required" json:"username"`
	Password 	string 	`gorm:"not null" json:"password"`
	Name		string	`gorm:"not null" valid:"required~your name is required" json:"name"`
	Email 		string 	`gorm:"not null;unique" valid:"required~your email is required, email~invalid email format" json:"email"`
	Date_Birth	string 	`gorm:"not null" valid:"required~your  birth  is required" json:"datebirth"`
	Gender		string	`gorm:"not null" valid:"required~your gender is required" json:"gender"`
}

//Menyimpan katasandi dalam atribut Password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FindUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}