<<<<<<< HEAD
# Build Stage
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o registerUser main.go

# Run Stage
FROM ubuntu:20.04
WORKDIR /app
COPY --from=builder /app/registerUser .
COPY --from=builder /app/.env .
EXPOSE 8081
CMD ["./registerUser"]