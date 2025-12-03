.PHONY: swagger run seed

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g main.go -o docs --parseDependency --parseInternal
	@echo "✅ Swagger documentation generated!"

# Run the application
run:
	@go run main.go

# Run seeder
seed:
	@go run main.go --seed

# Run seeder separately
seed-cmd:
	@go run cmd/seed.go

# Install dependencies
deps:
	@go mod download
	@go mod tidy

# Install swag CLI (if not installed)
install-swag:
	@go install github.com/swaggo/swag/cmd/swag@latest

