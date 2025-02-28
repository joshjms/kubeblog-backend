# Build stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
