# Setup Swagger Documentation

## Instalasi

1. Install Swagger CLI:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Atau menggunakan Makefile:
```bash
make install-swag
```

2. Install dependencies:
```bash
go mod tidy
```

## Generate Swagger Documentation

Setelah menambahkan atau mengubah annotations di controllers, jalankan:

```bash
swag init -g main.go -o docs --parseDependency --parseInternal
```

Atau menggunakan Makefile:
```bash
make swagger
```

## Menjalankan Aplikasi

1. Jalankan aplikasi:
```bash
go run main.go
```

Atau:
```bash
make run
```

2. Akses Swagger UI:
```
http://localhost:9001/swagger/index.html
```

## Menjalankan Seeder

1. Dengan flag:
```bash
go run main.go --seed
```

2. Atau menggunakan command terpisah:
```bash
go run cmd/seed.go
```

Atau:
```bash
make seed
```

## Struktur Swagger

- Swagger annotations ada di:
  - `main.go` - General API info
  - `app/http/controllers/AuthController.go` - Auth endpoints
  - `app/http/controllers/UserController.go` - User CRUD endpoints

- Generated docs ada di folder `docs/`

## Testing dengan Swagger UI

1. Buka http://localhost:9001/swagger/index.html
2. Klik "Authorize" button
3. Masukkan token JWT yang didapat dari `/api/auth/login`
4. Format: `Bearer <your_token>` atau langsung `<your_token>`
5. Klik "Authorize" untuk mengaktifkan token
6. Test semua endpoints langsung dari Swagger UI

