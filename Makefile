# Import docker Config
-include docker.env
export

# Extract the current directory name
APP_NAME := $(notdir $(patsubst %/,%,$(CURDIR)))
AFTER_FIRST_CMD = $(filter-out $@,$(MAKECMDGOALS))

# List Real Command
GET_CMD = go get $(AFTER_FIRST_CMD)
RUN_CMD = go run main.go http
TEST_CMD = go test -race -v -cover -covermode=atomic ./...
TIDY_CMD = go mod tidy
VERSION_CMD = git describe --tags
ifeq ($(filter-out sync,$(MAKECMDGOALS)),)
	SYNC_CMD = git pull
else
	SYNC_CMD = git fetch --all --tags --prune && git checkout tags/$(AFTER_FIRST_CMD)
endif
SETUP_CMD = cp sample.config.yaml config.yaml && cp sample.docker.env docker.env
DOCKER_UP_CMD = CONTAINER_NAME=$(APP_NAME)-app docker-compose --env-file docker.env up -d --build app
DOCKER_DOWN_CMD = CONTAINER_NAME=$(APP_NAME)-app docker-compose --env-file docker.env stop app
DOCKER_LOGS_CMD = docker logs -f $(APP_NAME)-app
DOCKER_STATS_CMD = docker stats $(APP_NAME)-app

# Update Command If OS Windows
ifeq ($(OS),Windows_NT)
endif

.PHONY: get run test tidy version sync setup up down logs stats dev sit prod

get:
	@${GET_CMD}

run:
	@$(RUN_CMD)

test:
	@$(TEST_CMD)

tidy:
	@$(TIDY_CMD)

version:
	@${VERSION_CMD}

sync:
	@${SYNC_CMD}

setup:
	@$(SETUP_CMD)

up:
	@$(DOCKER_UP_CMD)

down:
	@$(DOCKER_DOWN_CMD)

logs:
	@$(DOCKER_LOGS_CMD)

stats:
	@$(DOCKER_STATS_CMD)