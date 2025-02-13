using MySql.Data.MySqlClient;
using System;

namespace UpdatePayment.config
{
    public class Database
    {
        private readonly string _connectionString;

        public Database()
        {
            // Obtener configuración desde variables de entorno
            var host = Environment.GetEnvironmentVariable("DB_HOST");
            var user = Environment.GetEnvironmentVariable("DB_USER");
            var password = Environment.GetEnvironmentVariable("DB_PASSWORD");
            var database = Environment.GetEnvironmentVariable("DB_NAME");
            var port = Environment.GetEnvironmentVariable("DB_PORT") ?? "3306"; // Valor predeterminado

            // Construir la cadena de conexión
            _connectionString = $"Server={host};Port={port};Database={database};User={user};Password={password};";
        }

        public MySqlConnection GetConnection()
        {
            return new MySqlConnection(_connectionString);
        }
    }
}
