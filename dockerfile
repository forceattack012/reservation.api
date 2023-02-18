# Start with the official Golang image as the base
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main ./cmd

# Expose port 8080 for the application to listen on
EXPOSE 8080

# Start the application inside the container
CMD ["./main"]
