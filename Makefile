-include .env

up:
	docker compose build && docker compose up -d

logs:
	docker compose logs -f

down:
	docker compose down

go:
	docker exec -it $(MINIO_CONTROL_CONTAINER_HOST) /bin/sh