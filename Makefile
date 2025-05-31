COMPOSE_FILE = ./config/docker-compose.yaml
ARCHITECTURE = amd64
BUILD_DIR = $(ARCHITECTURE)
PLATFORM = windows
PROGRAMM_NAME = API.exe
CGO_ENABLED = 0
.PHONY: build up down stop logs


run:
	go run ./cmd

buildEXE:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(PLATFORM) GOARCH=$(ARCHITECTURE) go build -o ./$(BUILD_DIR)/$(PROGRAMM_NAME) ./cmd

build:
	docker-compose -f $(COMPOSE_FILE) build

up:
	docker-compose -f $(COMPOSE_FILE) up

down:
	docker-compose -f $(COMPOSE_FILE) down

logs:
	docker-compose -f $(COMPOSE_FILE) logs
	
clean:
	rm -rf $(BUILD_DIR)

all: build up