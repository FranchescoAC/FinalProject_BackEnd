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
