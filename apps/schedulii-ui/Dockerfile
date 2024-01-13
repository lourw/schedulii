# Notes on Dockerfile
# To build: `docker build -t frontend .`
# To run: `docker run -it --rm --init -p 8080:80 frontend`

# Install PNPM and node modules
FROM node:20-alpine as build-stage
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable && \
    apk add --update --no-cache python3 build-base # for some reason, node-gyp requires python

# Install node modules and build app
WORKDIR /app
COPY . .
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile && \
  npx nx build schedulii-ui

# Copy app to nginx
FROM nginx:stable-alpine
COPY ./apps/schedulii-ui/nginx.conf .
COPY --from=build-stage /app/dist/apps/schedulii-ui /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
