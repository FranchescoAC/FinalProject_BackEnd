using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using RegisterPayment.controller;
using RegisterPayment.config;

namespace RegisterPayment.routes
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