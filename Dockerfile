FROM node:18 as frontend-build
# Set up npm dependencies
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install

# Build assets for frontend
COPY ./frontend/. ./
RUN npm run build

FROM golang as backend-build
# Setup golang dependencies
WORKDIR /app/backend
COPY ./backend/go* ./
RUN go mod download && go mod verify

# Build deployment binary
COPY ./backend/. ./
RUN CGO_ENABLED=0 go build -o bin/schedulii ./src/main.go

FROM alpine:latest
# Configure for golang compatibility
RUN apk add gcompat && apk add --update curl

# Ensure server runs in production mode
ENV GIN_MODE=release

# Grab binaries and build assets from build stages
WORKDIR /app
COPY --from=frontend-build /app/build ./frontend/build
COPY --from=backend-build /app/backend/bin/. ./backend/bin/.
COPY --from=backend-build /app/backend/src/utils/credentials.json ./backend/src/utils/credentials.json

# Run application
EXPOSE 8080
WORKDIR /app/backend/bin
CMD ["./schedulii"]
