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

build: build_frontend build_backend

build_frontend:
	cd ${FRONTEND_DIR} && \
	npm run build

build_run_backend: build_backend run

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

ci_frontend: lint_frontend test_unit_frontend

# Requires ENV=test when calling make
ci_backend: lint_backend test_integration

# --- STATIC CODE ANALYSIS ---

lint_fix_frontend:
	cd ${FRONTEND_DIR} && \
	npm run lint-fix && \
	cd ..

lint_frontend:
	cd ${FRONTEND_DIR} && \
	npm run lint && \
	cd ..

lint_backend:
	cd ${BACKEND_DIR} && \
	golangci-lint run && \
	go vet ./... && \
	cd ..

# --- TESTING ---

test_unit_frontend:
	cd ${FRONTEND_DIR} && \
	npm run test -- --watchAll=false && \
	cd ..

test_unit_backend:
	cd ${BACKEND_DIR} && \
	go test ./... && \
	cd ..

test_integration: test_only	test_integration_setup test_unit_backend test_integration_teardown

test_integration_setup: test_only
	docker compose up -d

test_integration_teardown: test_only
	docker compose down --volumes

# When running make commands for testing, set flag to ENV=test
test_only:
ifneq ($(ENV),test)
	$(error ENV must be test)
endif
