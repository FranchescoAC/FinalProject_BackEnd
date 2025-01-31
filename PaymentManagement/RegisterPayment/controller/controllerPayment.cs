// Estructura y contenido de los archivos
// Payment Management/RegisterPayment/controller/controllerPayment.cs
using Microsoft.AspNetCore.Mvc;
using RegisterPayment.model;
using RegisterPayment.config;
using MySql.Data.MySqlClient;

namespace RegisterPayment.controller
{
    [ApiController]
    [Route("api/[controller]")]
    public class ControllerPayment : ControllerBase
    {
        private readonly Database _database;

        public ControllerPayment(Database database)
        {
            _database = database;
        }

        [HttpPost("register")]
        public IActionResult RegisterPayment([FromBody] Payment payment)
        {
            using var connection = _database.GetConnection();
            connection.Open();

            var query = "INSERT INTO payments (amount, description, date) VALUES (@amount, @description, @date);";
            using var command = new MySqlCommand(query, connection);
            command.Parameters.AddWithValue("@amount", payment.Amount);
            command.Parameters.AddWithValue("@description", payment.Description);
            command.Parameters.AddWithValue("@date", payment.Date);

            var result = command.ExecuteNonQuery();

            return result > 0 ? Ok("Payment registered successfully!") : StatusCode(500, "Error registering payment.");
        }
    }
}
