using Microsoft.AspNetCore.Builder;

namespace ListPayment.routes
{
    public static class RoutePayment
    {
        public static void RegisterPaymentRoutes(this WebApplication app)
        {
            app.MapControllers();
        }
    }
}
