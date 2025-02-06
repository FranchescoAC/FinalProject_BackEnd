# Usar una imagen base oficial de Go
FROM golang:1.18-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Download the missing jwt package
RUN go get github.com/golang-jwt/jwt/v4

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto 2002
EXPOSE 2002

# Comando para ejecutar la aplicación
CMD ["./main"]
