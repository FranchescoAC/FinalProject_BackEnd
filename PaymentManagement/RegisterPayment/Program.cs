using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using RegisterPayment.config;
using RegisterPayment.routes;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();

// Configure database connection
var connectionString = "Server=paymentdb.ckkft1kjllti.us-east-1.rds.amazonaws.com;Database=payment_management;User=admin;Password=Basepayment0806;";
builder.Services.AddSingleton(new Database(connectionString));

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
}

app.UseRouting();
app.MapControllers();
app.RegisterRoutes();

app.Run();