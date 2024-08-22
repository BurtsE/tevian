.PHONY: run ## Start application 
run:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} go run ./cmd
	
.PHONY: docker-up## Start application (requires running Compose services)
docker-up:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} docker-compose -f deploy/compose.yml up

.PHONY: docker-down
docker-down:
	@dotenv -f ./.env run -- env ${dev-env-vars} docker-compose -f deploy/compose.yml  down