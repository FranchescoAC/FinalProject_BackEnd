using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using DeletePayment.controller;
using DeletePayment.config;

namespace DeletePayment.routes
{
    public static class RoutePayment
    {
        public static void RegisterRoutes(this WebApplication app)
        {
            app.MapControllerRoute(
                name: "default",
                pattern: "{controller=ControllerPayment}/{action=DeletePayment}/{id?}");
        }
    }
}