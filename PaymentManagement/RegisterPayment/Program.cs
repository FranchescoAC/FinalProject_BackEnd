using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using RegisterPayment.config;
using RegisterPayment.routes;

var builder = WebApplication.CreateBuilder(args);

// ✅ Cargar la configuración desde variables de entorno
var port = Environment.GetEnvironmentVariable("PORT") ?? "7000"; // Puerto predeterminado: 7000

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
builder.Services.AddSingleton<Database>(); // Conexión a la base de datos

var app = builder.Build();

// ✅ Configurar el pipeline de la aplicación
if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
}

app.UseRouting();
app.UseCors("AllowAllOrigins"); // Habilitar CORS para acceso externo
app.MapControllers();
app.RegisterRoutes();

// ✅ Configurar la aplicación para escuchar en 0.0.0.0 y el puerto 7000
app.Urls.Add($"http://0.0.0.0:{port}");

Console.WriteLine($"✅ Server is running on port {port}");

app.Run();
