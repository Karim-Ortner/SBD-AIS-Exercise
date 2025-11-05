# todo
FROM golang:1.25 AS builder

# Set working directory
WORKDIR /app

# Copy project files into the container
COPY . .

# Run the build script (which outputs /app/ordersystem)
RUN chmod +x scripts/build-application.sh && \
    bash scripts/build-application.sh


FROM alpine:latest

# Create working directory
WORKDIR /app

# Copy compiled binary from builder stage
COPY --from=builder /app/ordersystem /app/ordersystem

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["/app/ordersystem"]