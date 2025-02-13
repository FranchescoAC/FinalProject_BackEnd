using Microsoft.AspNetCore.Mvc;
using ListPayment.config;
using ListPayment.model;
using MySql.Data.MySqlClient;
using System;
using System.Collections.Generic;

namespace ListPayment.controller
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

        [HttpGet("list")]
        public IActionResult ListPayments()
        {
            try
            {
                using var connection = _database.GetConnection();
                connection.Open();
                var command = new MySqlCommand("SELECT * FROM payments", connection);
                var reader = command.ExecuteReader();

                var payments = new List<Payment>();
                while (reader.Read())
                {
                    payments.Add(new Payment
                    {
                        Id = reader.GetInt32("id"),
                        Amount = reader.GetDecimal("amount"),
                        Description = reader.GetString("description"),
                        Date = reader.GetDateTime("date")
                    });
                }

                return Ok(payments);
            }
            catch (Exception ex)
            {
                return StatusCode(500, new { error = $"❌ Database error: {ex.Message}" });
            }
        }
    }
}
