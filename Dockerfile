# Step 1: Build the Go application
FROM golang:1.22-alpine AS build

# Set the current working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Set the working directory to the `cmd` folder and build the Go application
WORKDIR /app/cmd

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go run main.go

# Step 2: Create the runtime container
FROM alpine:3.18

# Set the current working directory in the container
WORKDIR /root/

# Copy the built binary from the build container
COPY --from=build /app/cmd/main .

# Expose the application port (replace with your actual port, e.g., 8080)
EXPOSE 8001

# Command to run the Go application
CMD ["./main"]
