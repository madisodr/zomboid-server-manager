APPLICATION_NAME := zomboid-server-manager
BUILD_VERSION := 0.0.1

DOCKER_USERNAME := winterthemute
DOCKER_DEFAULT_PLATFORM := linux/amd64
DOCKER_IMAGE := $(DOCKER_USERNAME)/$(APPLICATION_NAME)

BIN_DIRECTORY := ./bin
GO_CMD_DIR := ./cmd

export DOCKER_DEFAULT_PLATFORM

.PHONY: build run
build:
	go build -o $(BIN_DIRECTORY)/$(APPLICATION_NAME) $(GO_CMD_DIR)

run:
	$(BIN_DIRECTORY)/$(APPLICATION_NAME)

# Docker
.PHONY: docker-build publish
docker-build:
	docker build -t $(DOCKER_IMAGE) .
	
docker-run:
	docker run --env-file ./config/.env --rm -it $(DOCKER_IMAGE):latest

publish:
	docker tag $(DOCKER_USERNAME)/$(APPLICATION_NAME):latest $(DOCKER_USERNAME)/$(APPLICATION_NAME):$(BUILD_VERSION) && \
  docker push $(DOCKER_USERNAME)/$(APPLICATION_NAME):latest && \
	docker push $(DOCKER_USERNAME)/$(APPLICATION_NAME):$(BUILD_VERSION)