const mysql = require('mysql2/promise'); // Usa mysql2/promise
const dotenv = require('dotenv');

dotenv.config();

const connection = mysql.createPool({
    host: 'localhost',       // O la IP del servidor MySQL si no es local
    user: 'root',     // El usuario que creaste en MySQL (NO root, a menos que sea estrictamente necesario)
    password: 'root', // La contraseña del usuario
    database: 'bus_management', // El nombre de tu base de datos
    port: 3306,             // El puerto de MySQL (generalmente 3306)
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0
});

module.exports = connection;
