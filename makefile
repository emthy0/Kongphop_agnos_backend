# Define Docker Compose command
DOCKER_COMPOSE := docker-compose

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Available targets:"
	@echo "  test            - Run tests"
	@echo "  run             - Start the application locally with Docker Compose"
	@echo "  deploy          - Deploy the application to production"
.ONESHELL:
.PHONY: test
test:
	@echo "Running tests..."
	go test ./app/domain -cover -coverprofile=cover.out
	go tool cover -html=cover.out 

.PHONY: build
build: 
	@echo "building the application locally..."
	docker build ./app
.PHONY: run
run:
	@echo "Starting the application locally..."
	$(DOCKER_COMPOSE) up

