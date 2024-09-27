# ---- Stage 1: Build the Svelte frontend ----
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend


# Install dependencies and build the Svelte frontend
COPY frontend/package*.json ./
RUN npm install --include=dev

COPY frontend/ ./
RUN npm run build

# ---- Stage 2: Build the Go backend ----
FROM golang:1.22.1-alpine AS backend-builder

# Install required packages
RUN apk add --no-cache gcc musl-dev

WORKDIR /app/backend

# Copy frontend build to backend
COPY --from=frontend-builder /app/frontend/build /app/frontend/build

COPY backend/ ./
ENV CGO_ENABLED=1
RUN go mod download
RUN go build -o server main.go

# Expose a port for the Go server
EXPOSE 8080

# Start the Go server
CMD ["./server"]
    