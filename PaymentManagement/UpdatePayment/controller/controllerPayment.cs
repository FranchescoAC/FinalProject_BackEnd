using Microsoft.AspNetCore.Mvc;
using UpdatePayment.config;
using UpdatePayment.model;
using MySql.Data.MySqlClient;
using System;

namespace UpdatePayment.controller
{
    [Route("api/[controller]")]
    [ApiController]
    public class ControllerPayment : ControllerBase
    {
        private readonly Database _database;

        public ControllerPayment(Database database)
        {
            _database = database;
        }

        [HttpPut("update")]
        public IActionResult UpdatePayment([FromBody] Payment payment)
        {
            try
            {
                using var connection = _database.GetConnection();
                connection.Open();

                var command = new MySqlCommand(
                    "UPDATE payments SET amount = @amount, description = @description, date = @date WHERE id = @id",
                    connection
                );
                command.Parameters.AddWithValue("@amount", payment.Amount);
                command.Parameters.AddWithValue("@description", payment.Description);
                command.Parameters.AddWithValue("@date", payment.Date);
                command.Parameters.AddWithValue("@id", payment.Id);

                int rowsAffected = command.ExecuteNonQuery();

                if (rowsAffected == 0)
                {
                    return NotFound(new { error = "❌ Payment not found" });
                }

                return Ok(new { message = "✅ Payment updated successfully" });
            }
            catch (Exception ex)
            {
                return StatusCode(500, new { error = $"❌ Database error: {ex.Message}" });
            }
        }
    }
}
