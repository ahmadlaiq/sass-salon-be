package seeders

import (
	"log"
	"gorm.io/gorm"
)

// RunAllSeeders menjalankan semua seeder
func RunAllSeeders(db *gorm.DB) {
	log.Println("🌱 Memulai seeding database...")

	// Jalankan seeder users
	if err := SeedUsers(db); err != nil {
		log.Printf("❌ Error seeding users: %v\n", err)
	} else {
		log.Println("✅ Seeding users selesai")
	}

	log.Println("✅ Semua seeding selesai!")
}

