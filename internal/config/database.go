package config

import (
	"fmt"
	"log"
	"os"

	"kkp-backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database!", err)
	}

	DB = database
	log.Println("Database berhasil terkoneksi!")

	log.Println("Memulai migrasi database secara bertahap...")

	if err := DB.AutoMigrate(&models.User{}, &models.Bus{}, &models.Route{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tahap 1:", err)
	}

	if err := DB.AutoMigrate(&models.Seat{}, &models.Schedule{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tahap 2:", err)
	}

	if err := DB.AutoMigrate(&models.Booking{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tahap 3:", err)
	}

	if err := DB.AutoMigrate(&models.Payment{}, &models.Ticket{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tahap 4:", err)
	}

	log.Println("Migrasi tabel berhasil!")
}

func SeedUser() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	// Jika tabel user masih kosong, buat satu local user
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := models.User{
			UserID:      "U-001",
			FullName:    "Admin KKP",
			Email:       "admin@kkp.com",
			PhoneNumber: "081234567890",
			Password:    string(hashedPassword),
		}

		if err := DB.Create(&user).Error; err != nil {
			log.Fatal("Gagal membuat local user:", err)
		}
		log.Println("Seeder: Local user berhasil dibuat! (Email: admin@kkp.com | Pass: password123)")
	}
}
