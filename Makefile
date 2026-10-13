.PHONY: service-run

service-run:
	set -a && \
	. ./cmd/env/.env && \
	set +a && \
	export PORT=8080 && \
	go run ./cmd/API