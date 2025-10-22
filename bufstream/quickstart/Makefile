.phone: run
run: down up

.phony: down
down:
	docker compose down
	rm -rf data

.phony: up
up:
	docker compose up -d

.phony: stop
stop:
	docker compose stop

.phony: produce
produce:
	go run cmd/produce/main.go

.phony: configure-iceberg
configure-iceberg:



.phone: archive
archive:
	docker exec bufstream bufstream admin clean topics
