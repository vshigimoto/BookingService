grafana-run:
	docker run -d -p 3000:3000 --name=grafana grafana/grafana

fix-lint:
	golangci-lint run ./...

swag-v1-booking:
	swag init -g cmd/booking/main.go

swag-v1-auth:
	swag init -g internal/auth/server/http/router.go

swag-v1-user:
	swag init -g internal/user/server/http/router.go