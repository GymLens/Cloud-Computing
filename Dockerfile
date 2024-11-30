# Stage 1: Build the Go application
FROM golang:1.22 as builder

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application
RUN go build -o /bin/GymLens ./cmd/app

# Stage 2: Create a minimal Docker image
FROM gcr.io/distroless/base

WORKDIR /

# Copy the binary
COPY --from=builder /bin/GymLens /bin/GymLens

# Copy other necessary files (like config, if any)
COPY --from=builder /app/scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json /scripts/

# Set the command to run the application
CMD ["/bin/GymLens"]

# Expose port
EXPOSE 8080
