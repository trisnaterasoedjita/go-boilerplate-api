package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint   `gorm:"column=id"`
	Username     string `gorm:"column=username"`
	Name      string `gorm:"column=name"`
	Password  string `gorm:"-"` // ignore this field
	CreatedAt time.Time `gorm:"column=created_at"`
	UpdatedAt time.Time `gorm:"column=updated_at"`
	Db        *gorm.DB
}

// Checking whether user exist or not in database with
// provided email and password
func (u *User) IsUserExistByEmailPassword(email, password string) bool {
	var user User

	if !u.IsUserExistByEmail(email) {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// Checking whether user exist with email
func (u *User) IsUserExistByEmail(email string) bool {
	var count = 0
	u.Db.Where("username = ?", email).First(u).Count(&count)
	if count == 0 {
		return false
	}
	return true
}
