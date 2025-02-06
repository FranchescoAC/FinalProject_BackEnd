# Usamos una imagen base de Go
FROM golang:1.19

# Establecemos el directorio de trabajo
WORKDIR /app

# Download the missing jwt package
RUN go get github.com/golang-jwt/jwt/v4

# Copiamos el código al contenedor
COPY . .

# Instalamos las dependencias
RUN go mod tidy

# Exponemos el puerto en el que correrá el servicio
EXPOSE 2001

# Ejecutamos el servidor
CMD ["go", "run", "main.go"]
