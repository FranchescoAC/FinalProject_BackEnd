const TicketModel = require('../model/updateModel');

const updateTicket = async (req, res) => {
    try {
        const { id } = req.params;
        const { availabiliti, seat_number } = req.body; // Nombre corregido

        if (availabiliti === undefined && !seat_number) { // Nombre corregido y verifica si es undefined
            return res.status(400).json({ message: 'Debe proporcionar al menos un campo para actualizar (availabiliti o seat_number)' });
        }

        const updates = {};
        if (availabiliti !== undefined) { // Nombre corregido y verifica si es undefined
            updates.availabiliti = availabiliti; // Nombre corregido
        }
        if (seat_number) {
            updates.seat_number = seat_number;
        }

        const updated = await TicketModel.update(id, updates);

        if (!updated) {
            // ...
        }

        res.json({ message: 'Ticket updated successfully' });
    } catch (error) {
        // ...
    }
};

module.exports = { updateTicket };