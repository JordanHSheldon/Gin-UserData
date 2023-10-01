# Use an official Golang runtime as a parent image
FROM golang:1.21.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY ./src/go.mod .
COPY ./src/go.sum .
COPY ./src/main.go .

# Download and install Go dependencies
RUN go mod download

# Copy the local package files to the container's working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 on the container
EXPOSE 8080

# Command to run the application
CMD ["./main"]