BINARY_NAME=../bin/schedulii
BACKEND_DIR=./backend/src
FRONTEND_DIR=./frontend

include .env
export $(shell sed 's/=.*//' .env)

build_and_run: build run

build_run_backend: build_backend run

build: build_frontend build_backend

build_frontend:
	cd ${FRONTEND_DIR} && \
	npm run build

build_backend:
	cd ${BACKEND_DIR} && \
	go build -o ${BINARY_NAME} main.go

run:
	cd ${BACKEND_DIR} && \
	./${BINARY_NAME}

clean: 
	cd ${BACKEND_DIR} && \
	rm ${BINARY_NAME}

ci_frontend: 
	cd ${FRONTEND_DIR} && \
	npm run lint-fix && \
	npm test

ci_backend:
	cd ${BACKEND_DIR} && \
	golangci-lint run && \
	go vet ./... && \
	go test ./...
