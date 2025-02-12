const express = require('express');
const availabilityRoutes = require('./src/routes/availabilityRoutes');
const db = require('./src/config/db'); // Importa el módulo db
const app = express();
const PORT = process.env.PORT || 4004;

app.use(express.json());

// Middleware para manejar errores globalesS
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ message: 'Something went wrong!' });
});

// Rutas
app.use('/api/tickets', availabilityRoutes); // Prefijo para las rutas de tickets

// Manejo de la conexión a la base de datos (se elimina la prueba inicial)
async function startServer() {
    try {
        //const connection = await db.connect(); // Antes usabas db.connect()
        await db.getConnection(); // Ahora usa db.getConnection()
        console.log('Connected to MariaDB!');
        //connection.release(); // Ya no es necesario liberar la conexión aquí
        app.listen(PORT, '0.0.0.0', () => { // Escucha en todas las interfaces
            console.log(`Servidor corriendo en el puerto ${PORT}`);
        });
    } catch (error) {
        console.error('Failed to connect to MariaDB:', error);
        process.exit(1);
    }
}

startServer();