# ==============================================================================
# Main

runUser:
	echo "Starting user service"
	go run ./cmd/user/main.go

runAuth:
	echo "Starting Auth Service"
	go run ./cmd/auth/main.go

runBooking:
	echo "Starting booking service"
	go run ./cmd/booking/main.go

# ==============================================================================
# Tools commands

fix-lint:
	golangci-lint run ./...

swag-v1-booking:
	swag init -g cmd/booking/main.go


# ==============================================================================
# Go migrate postgresql

migrate_up:
	goose -dir migration/user postgres "user=postgres password=postgres port=5432 dbname=user sslmode=disable" up 20231207072610_add_user_table.sql
	goose -dir migration/booking postgres "user=postgres password=postgres port=5433 dbname=hotel sslmode=disable" up 20231207074148_add_booking_tabel.sql
	goose -dir migration/auth postgres "user=postgres password=postgres port=5434 dbname=auth sslmode=disable" up 20231207074229_add_auth_tabel.sql

migrate_down:
	goose -dir migration/user postgres "user=postgres password=postgres port=5432 dbname=user sslmode=disable" down 20231207072610_add_user_table.sql
	goose -dir migration/booking postgres "user=postgres password=postgres port=5433 dbname=hotel sslmode=disable" down 20231207074148_add_booking_tabel.sql
	goose -dir migration/auth postgres "user=postgres password=postgres port=5434 dbname=auth sslmode=disable" down 20231207074229_add_auth_tabel.sql


# ==============================================================================
# Docker compose commands
local:
	echo "Starting local environment"
	docker-compose -f docker/docker-compose.yml up --build