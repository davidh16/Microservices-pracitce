AUTH_BINARY=authentificationApp
CREATEUSER_BINARY=createUserApp
LOGGER_BINARY=loggerApp


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_logger build_createuser build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./authentification-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_createuser:
	@echo "Building auth binary..."
	cd ./createuser-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CREATEUSER_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	@echo "Building auth binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

