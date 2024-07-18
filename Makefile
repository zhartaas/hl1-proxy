up:
	docker-compose up

up-deattached:
	docker-compose up -d

# attach
logs:
	docker compose logs -f

build:
	docker-compose build

up-with-build:
	docker-compose up --build

down:
	docker-compose down

swagger:
	cd cmd/web && swag init

