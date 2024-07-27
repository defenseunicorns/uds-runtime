# Stage 1: Build Svelte frontend
FROM node:21-alpine AS frontend-builder
WORKDIR /app/ui
COPY ui/ ./
RUN npm ci && npm run build


# Stage 2: Build Go backend for amd64
FROM golang:1.22.4-alpine AS backend-builder-amd64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /app/ui/build ./ui/build
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /app/main-amd64 .

# Stage 2: Build Go backend for arm64
FROM golang:1.22.4-alpine AS backend-builder-arm64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /app/ui/build ./ui/build
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o /app/main-arm64 .

# Stage 3: Create the final image for amd64
FROM alpine:latest AS final-amd64
WORKDIR /app
COPY --from=backend-builder-amd64 /app/main-amd64 ./main
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]

# Stage 3: Create the final image for arm64
FROM alpine:latest AS final-arm64
WORKDIR /app
COPY --from=backend-builder-arm64 /app/main-arm64 ./main
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]
#
## Stage 2: Build Go backend
#FROM golang:1.22.4-alpine AS backend-builder
#WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod tidy
#
## Copy Go source code
#COPY . .
#
## Copy built frontend files into the appropriate directory for embedding
#COPY --from=frontend-builder /app/ui/build ./ui/build
#
## Build the Go binary
#RUN CGO_ENABLED=0 go build -o /app/main .
#
## Stage 3: Create the final image
#FROM alpine:latest
#WORKDIR /app
#COPY --from=backend-builder /app/main ./
#ENV PORT=8080
#EXPOSE 8080
#CMD ["./main"]
