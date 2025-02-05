const TicketModel = require('../model/registerModel');

const registerTicket = async (req, res) => {
    try {
        const { user_id, bus_id, seat_number, status } = req.body;

        if (!user_id || !bus_id || !seat_number || !status) {
            return res.status(400).json({ message: 'Faltan campos obligatorios' });
        }

        // Datos para crear el ticket, incluyendo user_id
        const newTicketData = { user_id, bus_id, seat_number, status };
        const createdTicket = await TicketModel.create(newTicketData); // Crea el ticket en la BD

        // Verifica si createdTicket tiene un ID (esto es crucial)
        if (!createdTicket || !createdTicket.id) {
            console.error("❌ Error al crear el ticket: ID no generado.");
            return res.status(500).json({ message: 'Error al registrar el ticket' });
        }

        res.status(201).json({ message: '🎫 Ticket registrado correctamente', ticket: createdTicket }); // Devuelve el ticket completo
    } catch (error) {
        console.error("❌ Error en registerTicket (catch):", error);
        res.status(500).json({ message: 'Error al registrar el ticket' });
    }
};

module.exports = { registerTicket };
