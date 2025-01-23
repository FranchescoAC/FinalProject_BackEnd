const mysql = require('mysql2/promise'); // Importante: /promise

const pool = mysql.createPool({
    host: 'localhost', // O la dirección de tu servidor MariaDB
    user: 'root',      // Tu usuario de MariaDB
    password: 'root',  // Tu contraseña de MariaDB
    database: 'ticket_management', // El nombre de tu base de datos
    connectionLimit: 5, // Número máximo de conexiones en el pool
    port: 3307          // El puerto de MariaDB (usualmente 3306 o 3307)
});

async function getConnection() {
    try {
        const connection = await pool.getConnection();
        return connection;
    } catch (error) {
        console.error('Error al conectar a MariaDB:', error);
        throw error; // Importante: Re-lanza el error
    }
}

module.exports = { getConnection };