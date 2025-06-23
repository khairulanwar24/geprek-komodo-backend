package models

import (
	"ayam-geprek-backend/config"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"uniqueIndex;not null" json:"username"`
	Password  string     `json:"password"`
	Nama      string     `json:"nama"`
	Email     string     `json:"email"`
	NoHp      string     `json:"no_hp"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func RegisterUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return config.DB.Create(user).Error
}

// GetUserByUsername mencari user berdasarkan username
func GetUserByUsername(username string) (*User, error) {
	var user User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// CheckPassword memeriksa password cocok dengan hash
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
