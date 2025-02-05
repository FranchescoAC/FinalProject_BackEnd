# Usar una imagen base de Go
FROM golang:1.18-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod tidy

# Copiar todo el código fuente del proyecto al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto en el que la aplicación estará escuchando (puedes cambiar el puerto si es necesario)
EXPOSE 8083

# Ejecutar la aplicación cuando el contenedor inicie
CMD ["./main"]
