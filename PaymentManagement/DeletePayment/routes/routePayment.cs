using Microsoft.AspNetCore.Builder;
using DeletePayment.controller;

namespace DeletePayment.routes
{
    public static class RoutePayment // Puedes cambiar el nombre de la clase
    {
        // ✅ Nombre del método corregido para que coincida con la llamada en Program.cs
        public static void RegisterPaymentRoutes(this WebApplication app)
        {
            app.MapControllerRoute(
                name: "default",
                pattern: "{controller=ControllerPayment}/{action=DeletePayment}/{id?}");
        }
    }
}