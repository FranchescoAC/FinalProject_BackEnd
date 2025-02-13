using MySql.Data.MySqlClient;
using System;

namespace DeletePayment.config
{
    public class MySqlDatabase // ✅ Renombrar la clase para evitar conflictos
    {
        private readonly string _connectionString;

        public MySqlDatabase()
        {
            // Obtener configuración desde variables de entorno
            var host = Environment.GetEnvironmentVariable("DB_HOST");
            var user = Environment.GetEnvironmentVariable("DB_USER");
            var password = Environment.GetEnvironmentVariable("DB_PASSWORD");
            var database = Environment.GetEnvironmentVariable("DB_NAME");
            var port = Environment.GetEnvironmentVariable("DB_PORT") ?? "3306"; // Valor predeterminado

            // Construir la cadena de conexión
            _connectionString = $"Server={host};Port={port};Database={database};User={user};Password={password};";

            // ✅ Imprimir la cadena de conexión para verificarla
            Console.WriteLine($"Connection string: {_connectionString}");
        }

        public MySqlConnection GetConnection()
        {
            return new MySqlConnection(_connectionString);
        }
    }
}