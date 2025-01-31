using Microsoft.AspNetCore.Mvc;
using DeletePayment.config;
using MySql.Data.MySqlClient;

namespace DeletePayment.controller
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

        [HttpDelete("delete/{id}")]
        public IActionResult DeletePayment(int id)
        {
            try
            {
                using var connection = _database.GetConnection();
                connection.Open();

                var command = new MySqlCommand("DELETE FROM payments WHERE id = @id", connection);
                command.Parameters.AddWithValue("@id", id);

                int rowsAffected = command.ExecuteNonQuery();

                if (rowsAffected == 0)
                {
                    return NotFound(new { error = "Payment not found" });
                }

                return Ok(new { message = "Payment deleted successfully" });
            }
            catch (Exception ex)
            {
                return StatusCode(500, new { error = ex.Message });
            }
        }
    }
}
