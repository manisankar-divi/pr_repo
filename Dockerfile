# Use the official Go image to build and run the application
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp main.go

# Command to run the application
CMD ["./myapp"]
