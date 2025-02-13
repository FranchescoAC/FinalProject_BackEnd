require('dotenv').config(); // Importa dotenv para leer las variables de entorno
const mysql = require('mysql2/promise');

const pool = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_NAME,
    connectionLimit: 5,
    port: process.env.DB_PORT,
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