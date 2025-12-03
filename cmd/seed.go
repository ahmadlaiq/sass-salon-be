package main

import (
	"log"
	"time"

	"github.com/spf13/viper"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"gin-sass-salon/config"
	"gin-sass-salon/database/seeders"
)

func main() {
	// 1. Muat Konfigurasi (.env)
	config.LoadConfig()

	// 2. Koneksi ke Database PostgreSQL
	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		DisableAutomaticPing: false,
	})

	if err != nil {
		log.Fatalf("❌ Gagal koneksi ke database PostgreSQL: %v", err)
	}

	// Konfigurasi Pool Koneksi
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Koneksi Database PostgreSQL berhasil!")

	// 3. Jalankan seeder
	seeders.RunAllSeeders(db)

	log.Println("✅ Seeding selesai!")
}

