const express = require('express');
const updateRoutes = require('./src/routes/updateRoutes');
const db = require('./src/config/db');
const app = express();
const PORT = process.env.PORT || 4001;

app.use(express.json()); // Middleware para parsear JSON

// Middleware para manejar errores globales
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ message: 'Something went wrong!' });
});

// Rutas
app.use('/api/tickets', updateRoutes);

async function startServer() {
    try {
        await db.getConnection();
        console.log('Connected to MariaDB!');
        app.listen(PORT, '0.0.0.0', () => {
            console.log(`Servidor corriendo en el puerto ${PORT}`);
        });
    } catch (error) {
        console.error('Failed to connect to MariaDB:', error);
        process.exit(1);
    }
}

startServer();