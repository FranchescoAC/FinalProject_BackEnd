const express = require('express');
const bodyParser = require('body-parser');
const ticketRoutes = require('./src/routes/registerRoutes');
const db = require('./src/config/db');
const app = express();
const PORT = process.env.PORT || 4002;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use('/api/tickets', ticketRoutes);

async function startServer() {
    try {
        await db.getConnection();
        console.log('✅ Conectado a la base de datos.');
        app.listen(PORT, '0.0.0.0', () => { // Escucha en todas las interfaces
            console.log(` Servidor corriendo en http://localhost:${PORT}`);
        });
    } catch (error) {
        console.error('❌ Error conectando a la base de datos:', error);
        process.exit(1);
    }
}

app.use((err, req, res, next) => {
    console.error('❌ Error global:', err.stack);
    res.status(500).json({ message: 'Algo salió mal en el servidor. Intenta nuevamente más tarde.' });
});

startServer();