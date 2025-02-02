# Usamos una imagen base de Go
FROM golang:1.19

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos el código al contenedor
COPY . .

# Instalamos las dependencias
RUN go mod tidy

# Exponemos el puerto en el que correrá el servicio
EXPOSE 2001

# Ejecutamos el servidor
CMD ["go", "run", "main.go"]
