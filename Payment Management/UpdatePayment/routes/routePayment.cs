using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using UpdatePayment.controller;
using UpdatePayment.config;

namespace UpdatePayment.routes
{
    public static class RoutePayment
    {
        public static void RegisterRoutes(this WebApplication app)
        {
            app.MapControllerRoute(
                name: "default",
                pattern: "{controller=ControllerPayment}/{action=RegisterPayment}/{id?}");
        }
    }
}