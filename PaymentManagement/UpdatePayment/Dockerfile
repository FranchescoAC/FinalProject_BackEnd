# Usar la imagen oficial de .NET SDK para compilar y publicar la aplicación
FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /app

# Copiar archivos del proyecto y restaurar las dependencias
COPY *.csproj ./
RUN dotnet restore

# Copiar el resto de la aplicación y compilarla en modo release
COPY . ./
RUN dotnet publish -c Release -o /publish

# Usar la imagen de .NET runtime para ejecutar la aplicación
FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS runtime
WORKDIR /app

# Copiar la aplicación publicada desde la etapa anterior
COPY --from=build /publish .

# Exponer el puerto definido en el .env
ENV PORT=7002
EXPOSE ${PORT}

# Comando para ejecutar la aplicación
ENTRYPOINT ["dotnet", "UpdatePayment.dll"]
