package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(100);unique"`
	Image     []byte
	FirstName string `gorm:"type:varchar(20)"`
	LastName  string `gorm:"type:varchar(20)"`
	Title     string `gorm:"type:varchar(5)"`
	Password  string `gorm:"type:varchar(100)"`
	Role      string `gorm:"type:varchar(100)"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	// Hash the password before saving the user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
func (user *User) CheckPassword(password string) bool {
	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
