# Usar una imagen base de Go
FROM golang:1.18-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./

# Descargar todas las dependencias especificadas en go.mod
RUN go mod download

# Download the missing jwt package
RUN go get github.com/golang-jwt/jwt/v4

# Copiar todo el código fuente del proyecto al contenedor
COPY . .

# Descargar las dependencias adicionales y limpiar el módulo (si es necesario)
RUN go mod tidy

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto en el que la aplicación estará escuchando
EXPOSE 2004

# Ejecutar la aplicación cuando el contenedor inicie
CMD ["./main"]
