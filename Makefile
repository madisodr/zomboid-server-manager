APP_NAME = zomboid-server-manager
IMAGE_NAME = winterthemute/$(APP_NAME)
DOCKERFILE_PATH = .
SHELL := /bin/bash

.PHONY: all build build-docker

all: build build-docker

build:
	go build -o ./app/$(APP_NAME) ./cmd

run:
	source config.env && ./app/$(APP_NAME)

build-docker:
	docker build -t $(IMAGE_NAME) $(DOCKERFILE_PATH)