package config

import (
	"log"
	"fmt" // Tambahkan import fmt untuk string formatting
	"github.com/spf13/viper"
)

// DSN - Data Source Name untuk koneksi GORM (PostgreSQL Format)
func DSN() string {
	// Format DSN PostgreSQL: "host=... user=... password=... dbname=... port=... sslmode=disable TimeZone=Asia/Jakarta"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)
	return dsn
}

// LoadConfig (Isi tetap sama seperti sebelumnya)
func LoadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".") 

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error membaca file konfigurasi:", err)
	}

	log.Println("Konfigurasi (.env) berhasil dimuat.")
}

// JWTSecret mengembalikan JWT secret key dari environment variable
func JWTSecret() string {
	secret := viper.GetString("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key-change-this-in-production"
	}
	return secret
}