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
=======
# Etapa de construcción
FROM golang:1.21-alpine AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del proyecto
COPY . .

# Descargar las dependencias y compilar el binario
RUN go mod tidy && go build -o app

# Etapa final (imagen ligera)
FROM alpine:latest

# Establecer el directorio de trabajo
WORKDIR /root/

# Copiar el binario desde la imagen anterior
COPY --from=builder /app/app .

# Exponer el puerto en el que escucha la aplicación
EXPOSE 8081

# Ejecutar la aplicación
CMD ["./main"]
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
