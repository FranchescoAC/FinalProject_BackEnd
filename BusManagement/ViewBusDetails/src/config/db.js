const mysql = require('mysql2/promise');

const connection = mysql.createPool({
    host: 'busmanagement.ckkft1kjllti.us-east-1.rds.amazonaws.com', // RDS endpoint
    user: 'root', // RDS user
    password: 'root12345678', // RDS password
    database: 'bus_management', // Database name
    port: 3306, // Default port for MySQL
    waitForConnections: true, // Allow requests to wait if no connections are available
    connectionLimit: 10, // Maximum number of simultaneous connections
    queueLimit: 0 // No limit on the number of queued requests
});

module.exports = connection;
