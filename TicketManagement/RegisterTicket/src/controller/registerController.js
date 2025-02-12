const TicketModel = require('../model/registerModel');

const registerTicket = async (req, res) => {
    try {
        const { name_bus, seat_number } = req.body;

        if (!name_bus || !seat_number) {
            return res.status(400).json({ message: 'Faltan campos obligatorios' });
        }

        const newTicketData = { name_bus, seat_number };
        const createdTicket = await TicketModel.create(newTicketData);

        if (!createdTicket || !createdTicket.id) {
            console.error("❌ Error al crear el ticket: ID no generado.");
            return res.status(500).json({ message: 'Error al registrar el ticket' });
        }

        res.status(201).json({ message: ' Ticket registrado correctamente', ticket: createdTicket });
    } catch (error) {
        console.error("❌ Error en registerTicket (catch):", error);
        res.status(500).json({ message: 'Error al registrar el ticket' });
    }
};

module.exports = { registerTicket };