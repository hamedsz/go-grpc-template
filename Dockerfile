FROM golang:1.16.5 as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 6565
# Start app
CMD reflex -g '*.go' go run ./commands/server/server.go --start-service