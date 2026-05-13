APP=starter-api
COMPOSE=docker compose -f deployment/docker-compose.yml

build:
	$(COMPOSE) --env-file .env --profile dev build

up:
	$(COMPOSE) --env-file .env --profile dev up -d

down:
	$(COMPOSE) --env-file .env --profile dev down

logs:
	$(COMPOSE) --env-file .env --profile dev logs -f

restart:
	$(COMPOSE) --env-file .env --profile dev restart

prod-build:
	docker build -f deployment/docker/app/Dockerfile \
		--target prod \
		-t $(APP):latest .

prod-run:
	docker run -p 8080:8080 --env-file .env $(APP):latest