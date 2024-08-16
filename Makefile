.PHONY: run ## Start application 
run:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} go run ./cmd