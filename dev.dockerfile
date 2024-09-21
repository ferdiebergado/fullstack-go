# Start with the official Go image
FROM golang:1.22.6-bookworm

ARG APP_PORT

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Command to run Air
CMD ["air"]

# Expose the application port
EXPOSE ${APP_PORT}

# Mount your source files at /app when running the container
