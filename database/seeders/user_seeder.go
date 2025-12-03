package seeders

import (
	"log"
	"gorm.io/gorm"
	"gin-sass-salon/app/models"
)

// SeedUsers mengisi database dengan data user contoh
func SeedUsers(db *gorm.DB) error {
	// Data users yang akan di-seed
	users := []models.User{
		{
			Name:     "Admin",
			Email:    "admin@example.com",
			Password: "password123",
		},
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: "password123",
		},
		{
			Name:     "Bob Johnson",
			Email:    "bob@example.com",
			Password: "password123",
		},
		{
			Name:     "Alice Williams",
			Email:    "alice@example.com",
			Password: "password123",
		},
	}

	// Hash password dan simpan ke database
	for i := range users {
		// Cek apakah user sudah ada
		var existingUser models.User
		if err := db.Where("email = ?", users[i].Email).First(&existingUser).Error; err == nil {
			log.Printf("User dengan email %s sudah ada, dilewati\n", users[i].Email)
			continue
		}

		// Hash password
		if err := users[i].HashPassword(); err != nil {
			log.Printf("Gagal hash password untuk %s: %v\n", users[i].Email, err)
			continue
		}

		// Simpan ke database
		if err := db.Create(&users[i]).Error; err != nil {
			log.Printf("Gagal membuat user %s: %v\n", users[i].Email, err)
			continue
		}

		log.Printf("✅ User berhasil dibuat: %s (%s)\n", users[i].Name, users[i].Email)
	}

	return nil
}

