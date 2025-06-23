package models

import (
	"ayam-geprek-backend/config"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func SeedUsers() {
	users := []User{
		{
			Username: "admin",
			Password: "admin1234",
			Nama:     "Admin Utama",
			Email:    "admin@example.com",
			NoHp:     "081234567890",
			Role:     "admin",
		},
		{
			Username: "kasir",
			Password: "kasir123",
			Nama:     "Kasir 1",
			Email:    "kasir@example.com",
			NoHp:     "081111111111",
			Role:     "kasir",
		},
		{
			Username: "manager",
			Password: "manager123",
			Nama:     "Manager Outlet",
			Email:    "manager@example.com",
			NoHp:     "082222222222",
			Role:     "manager",
		},
	}

	for _, u := range users {
		var existing User
		result := config.DB.Where("username = ?", u.Username).First(&existing)

		if result.RowsAffected == 0 {
			// User belum ada, buat baru
			err := RegisterUser(&u)
			if err != nil {
				fmt.Printf("❌ Gagal membuat user %s: %v\n", u.Username, err)
			} else {
				fmt.Printf("✅ User %s berhasil dibuat!\n", u.Username)
			}
		} else {
			// User sudah ada, update password
			hashed, _ := HashPassword(u.Password)
			existing.Password = hashed
			config.DB.Save(&existing)
			// fmt.Printf("🔁 Password user %s di-reset!\n", u.Username)
		}
	}
}
