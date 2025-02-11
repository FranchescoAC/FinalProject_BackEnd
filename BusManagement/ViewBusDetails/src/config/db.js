const mysql = require('mysql2/promise');
const dotenv = require('dotenv');

dotenv.config();

const connection = mysql.createPool({
    host: process.env.DB_HOST || 'busmanagement.ckkft1kjllti.us-east-1.rds.amazonaws.com', // Aquí usas el endpoint de RDS
    user: process.env.DB_USER || 'admin', // Usa tu usuario de RDS
    password: process.env.DB_PASSWORD || 'tu-contraseña', // La contraseña de RDS
    database: process.env.DB_NAME || 'bus_management', // Nombre de tu base de datos
    port: process.env.DB_PORT || 3306, // Puerto 3306 por defecto para MySQL
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0
});

module.exports = connection;
