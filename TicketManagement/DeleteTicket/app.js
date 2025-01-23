const express = require('express');
const registerRoutes = require('../DeleteTicket/src/routes/deleteRoutes');
const db = require('../DeleteTicket/src/config/db'); // Importa el módulo db
const app = express();
const PORT = process.env.PORT || 4003;

app.use(express.json());

// Middleware para manejar errores globalesS
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ message: 'Something went wrong!' });
});

// Rutas
app.use('/api/tickets', registerRoutes); // Prefijo para las rutas de tickets

// Manejo de la conexión a la base de datos (se elimina la prueba inicial)
async function startServer() {
    try {
        //const connection = await db.connect(); // Antes usabas db.connect()
        await db.getConnection(); // Ahora usa db.getConnection()
        console.log('Connected to MariaDB!');
        //connection.release(); // Ya no es necesario liberar la conexión aquí
        app.listen(PORT, () => {
            console.log(`Server is running on http://localhost:${PORT}`);
        });
    } catch (error) {
        console.error('Failed to connect to MariaDB:', error);
        process.exit(1);
    }
}

startServer();