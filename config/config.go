package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// ✅ Load file .env dari folder app/
	err := godotenv.Load("app/.env")
	if err != nil {
		log.Println("⚠️ Gagal load app/.env, menggunakan environment default")
	}

	// ✅ Ambil isi dari variabel DATABASE_URL
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL belum diset di .env atau environment")
	}

	// ✅ Connect ke database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("✅ Berhasil connect ke database PostgreSQL")
}
