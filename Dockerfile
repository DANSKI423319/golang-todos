# Use the official Golang image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /application

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Use Air for hot reloading
CMD ["air", "-c", ".air.toml"]
