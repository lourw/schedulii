# Install rust
FROM rust:alpine as build-stage

ARG LOCAL_PATH=./apps/schedulii_api
ENV APP_LOCAL_PATH=$LOCAL_PATH

# Copy over Cargo files and build a dummy /src/main.rs
# This is done so we can cache our dependencies and speed up subsequent builds
WORKDIR /app
RUN apk add build-base && \
mkdir ./src && \
echo 'fn main() { println!("Dummy!"); }' > ./src/main.rs

COPY ${LOCAL_PATH}/Cargo.toml .
RUN cargo build --release && \
rm -rf ./src

# Copy application and build release
COPY ${LOCAL_PATH}/src ./src
RUN touch src/main.rs && \
cargo build --release

FROM alpine:latest
WORKDIR /app
COPY --from=build-stage /app/target/release/. .
EXPOSE 9000
CMD ["./schedulii_api"]
#
# Notes on Dockerfile
# To build: `docker build -t backend .`
# To run: `docker run -it --rm --init -p 9000:9000 backend`
