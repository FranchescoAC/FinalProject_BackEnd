using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.EntityFrameworkCore; // Importa Entity Framework Core
using DeletePayment.config; // Importa tu configuración
using DeletePayment.routes; // Importa tus rutas
using System;

var builder = WebApplication.CreateBuilder(args);

// ✅ Cargar configuración de entorno
var port = Environment.GetEnvironmentVariable("PORT") ?? "7003";

// ✅ Configurar CORS para permitir solicitudes desde cualquier origen
builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAllOrigins",
        policy => policy.AllowAnyOrigin()
            .AllowAnyMethod()
            .AllowAnyHeader());
});

// ✅ Agregar servicios al contenedor
builder.Services.AddControllers();

// ✅ Inyección de dependencia para tu DbContext (reemplaza MySqlDatabase con el nombre real)
builder.Services.AddDbContext<MySqlDatabase>(options => 
    options.UseMySql(builder.Configuration.GetConnectionString("YourConnectionString"))); // Usa MySql

// ✅ Inyección de dependencia para tu clase de base de datos (renombrada a MySqlDatabase)
builder.Services.AddSingleton<MySqlDatabase>();

var app = builder.Build();

// ✅ Configurar el pipeline de la aplicación
if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
}

app.UseRouting();
app.UseCors("AllowAllOrigins"); // ✅ Habilitar CORS para acceso externo
app.UseAuthorization(); // Si usas autorización

app.MapControllers();

// ✅ Llamar al método de extensión RegisterPaymentRoutes para registrar las rutas
app.RegisterPaymentRoutes(); // Debe coincidir el nombre del método en routePayment.cs

// ✅ Configurar la aplicación para escuchar en 0.0.0.0 y el puerto 7003
app.Urls.Add($"http://0.0.0.0:{port}");

Console.WriteLine($"✅ Server is running on port {port}");

app.Run();