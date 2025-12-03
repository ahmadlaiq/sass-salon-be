# Gin SASS Salon API

API untuk aplikasi SASS Salon yang dibangun menggunakan Go (Golang) dengan framework Gin, GORM (PostgreSQL), dan JWT Authentication.

## 🚀 Fitur

- **Framework**: [Gin Web Framework](https://gin-gonic.com/)
- **Database**: PostgreSQL dengan [GORM](https://gorm.io/)
- **Authentication**: JWT (JSON Web Token)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Documentation**: Swagger UI (via [swaggo](https://github.com/swaggo/swag))
- **Migration**: Auto Migration dengan GORM
- **Seeding**: Database seeder support

## 🛠️ Prasyarat

Pastikan Anda telah menginstal:

- [Go](https://go.dev/) (versi 1.24 atau lebih baru)
- [PostgreSQL](https://www.postgresql.org/)
- [Make](https://www.gnu.org/software/make/) (Opsional, untuk menjalankan perintah Makefile)

## 📦 Instalasi

1. **Clone repositori ini**
   ```bash
   git clone https://github.com/ahmadlaiq/sass-salon-be.git
   cd sass-salon-be
   ```

2. **Setup Environment Variables**
   Salin file `.env.example` ke `.env` dan sesuaikan konfigurasinya (terutama koneksi database).
   ```bash
   cp .env.example .env
   ```
   Edit `.env`:
   ```env
   APP_PORT=9001
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=password
   DB_NAME=sass_salon
   DB_PORT=5432
   JWT_SECRET=rahasia_negara
   ```

3. **Install Dependencies**
   ```bash
   go mod download
   ```

## 🏃‍♂️ Menjalankan Aplikasi

### Menggunakan Makefile

- **Jalankan Server**:
  ```bash
  make run
  ```

- **Generate Swagger Docs**:
  ```bash
  make swagger
  ```

- **Run Seeder**:
  ```bash
  make seed
  ```

### Manual

- **Jalankan Server**:
  ```bash
  go run main.go
  ```

- **Run Seeder**:
  ```bash
  go run main.go --seed
  ```

## 📚 Dokumentasi API

Setelah server berjalan, Anda dapat mengakses dokumentasi API melalui Swagger UI:

- **Swagger UI**: [http://localhost:9001/swagger/index.html](http://localhost:9001/swagger/index.html)
- **API Docs JSON**: [http://localhost:9001/api/docs/doc.json](http://localhost:9001/api/docs/doc.json)

## 🛣️ Endpoints Utama

### Auth
- `POST /api/auth/register` - Registrasi user baru
- `POST /api/auth/login` - Login dan dapatkan JWT

### Users (Protected)
- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create user
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

## 🤝 Kontribusi

Silakan buat Pull Request atau Issue jika menemukan bug atau ingin menambahkan fitur.

## 📄 Lisensi

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
