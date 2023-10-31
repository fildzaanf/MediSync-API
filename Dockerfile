# syntax=docker/dockerfile
FROM golang:1.21.0-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.app

EXPOSE 8000

# Run
CMD ["/app/main.app"]