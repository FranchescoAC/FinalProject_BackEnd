const TicketModel = require('../model/registerModel');

const registerTicket = async (req, res) => {
    try {
        const { name_bus, seat_number, availabiliti } = req.body; // Nombre corregido

        if (!name_bus || !seat_number || availabiliti === undefined) { // Nombre corregido y verifica si es undefined
            return res.status(400).json({ message: 'Faltan campos obligatorios: name_bus, seat_number y availabiliti' });
        }

        const newTicketData = { name_bus, seat_number, availabiliti }; // Nombre corregido
        const createdTicket = await TicketModel.create(newTicketData);

        if (!createdTicket || !createdTicket.id) {
            // ...
        }

        res.status(201).json({ message: 'Ticket registrado correctamente', ticket: createdTicket });
    } catch (error) {
        // ...
    }
};

module.exports = { registerTicket };