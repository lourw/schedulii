# ENV can be empty or 'test'
ENV=prod
-include .env
-include .env.$(ENV)
export

BINARY_NAME=../bin/schedulii
BACKEND_DIR=./backend/src
FRONTEND_DIR=./frontend

# --- BUILD COMMANDS ---

build_and_run: build run

build_run_backend: build_backend run

build: build_frontend build_backend

build_frontend:
	cd ${FRONTEND_DIR} && \
	npm run build

build_backend:
	cd ${BACKEND_DIR} && \
	go build -o ${BINARY_NAME} main.go wire_gen.go app.go

run:
	cd ${BACKEND_DIR} && \
	./${BINARY_NAME}

clean: 
	cd ${BACKEND_DIR} && \
	rm ${BINARY_NAME}

# --- CONTINUOUS INTEGRATION ---

ci_frontend: 
	cd ${FRONTEND_DIR} && \
	npm run lint-fix && \
	npm test && \
	cd ..

ci_backend:
	cd ${BACKEND_DIR} && \
	golangci-lint run && \
	go vet ./... && \
	go test ./... && \
	cd ..

# --- TESTING ---

test_integration: test_only	test_integration_setup ci_backend test_integration_teardown

test_integration_setup: test_only
	docker compose up -d

test_integration_teardown: test_only
	docker compose down --volumes

# When running make commands for testing, set flag to ENV=test
test_only:
ifneq ($(ENV),test)
	$(error ENV must be test)
endif
