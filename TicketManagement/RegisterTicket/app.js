const express = require('express');
const TicketModel = require('./src/model/registerModel'); // Importar el modelo correcto
const db = require('./src/config/db');
const app = express();
const PORT = process.env.PORT || 4002;

app.use(express.json());

// Endpoint para recibir el Webhook de RegisterUser
app.post('/api/tickets/register-ticket', async (req, res) => {
    const { user_id, username, email } = req.body;

    if (!user_id || !username || !email) {
        return res.status(400).json({ message: '❌ Datos inválidos en el Webhook' });
    }

    try {
        console.log(`📩 Webhook recibido: Usuario ${username} (${email}) con ID ${user_id} registrado.`);

        // Crear un ticket con el user_id recibido
        const ticketData = {
            user_id: user_id,
            bus_id: 1, // Se puede cambiar dinámicamente
            seat_number: Math.floor(Math.random() * 50) + 1, // Número de asiento aleatorio
            status: "reserved"
        };

        // Guardar en la base de datos
        const ticket = await TicketModel.create(ticketData);

        res.status(200).json({
            message: `🎫 Ticket para el usuario ${username} creado exitosamente`,
            ticket
        });
    } catch (error) {
        console.error('❌ Error procesando el Webhook:', error);
        res.status(500).json({ message: 'Error procesando el Webhook' });
    }
});

// Conectar a la base de datos y arrancar el servidor
async function startServer() {
    try {
        await db.getConnection();
        console.log('✅ Conectado a la base de datos.');
        app.listen(PORT, () => {
            console.log(`🚀 Servidor corriendo en http://localhost:${PORT}`);
        });
    } catch (error) {
        console.error('❌ Error conectando a la base de datos:', error);
        process.exit(1);
    }
}

// Middleware para manejar errores (debe ir después de las rutas)
app.use((err, req, res, next) => {
    console.error('❌ Error global:', err.stack);  // Muestra el error completo en los logs
    res.status(500).json({ message: 'Algo salió mal en el servidor. Intenta nuevamente más tarde.' });
});

startServer();
