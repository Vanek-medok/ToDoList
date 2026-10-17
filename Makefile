.PHONY: service-run

service-run:
	@if [ -f ./cmd/env/.env ]; then \
		set -a && . ./cmd/env/.env && set +a; \
	else \
		echo "Предупреждение: .env файл не найден, используются внешние переменные"; \
	fi
	go run ./cmd/API/main.go

migrate-up:
	@if [ -f ./cmd/env/.env ]; then \
		set -a && . ./cmd/env/.env && set +a; \
	else \
		echo "Предупреждение: .env файл не найден, используются внешние переменные"; \
	fi
	migrate -path ./migrations -database "$$DATABASE_URL" up