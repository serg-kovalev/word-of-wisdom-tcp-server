# Use the official Go image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the server binary
RUN go build -o word-of-wisdom-server cmd/main.go

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy only the necessary files from the builder image
COPY --from=builder /app/word-of-wisdom-server /app/word-of-wisdom-server

# Expose the port that the server will run on
EXPOSE 8080

# Run the server
CMD ["./word-of-wisdom-server"]
