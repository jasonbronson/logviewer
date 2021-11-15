APP_NAME   := api
APP_DIR := ./*.go
DIST_DIR = ./dist

setup:
	go mod download

test:
	go test -parallel=6 -failfast -cover ./...
	
build: ## Build go binary
	go build -o ${DIST_DIR}/${APP_NAME} ${APP_DIR}

buildcron: 
	go build -o ${DIST_DIR}/cron ./transport/cron.go

local: 
	docker-compose up