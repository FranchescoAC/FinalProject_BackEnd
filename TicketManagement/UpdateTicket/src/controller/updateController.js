const TicketModel = require('../model/updateModel');

const updateTicket = async (req, res) => {
    try {
        const { id } = req.params;
        const { status, seat_number } = req.body;

        const updated = await TicketModel.update(id, { status, seat_number });
        if (!updated) {
            return res.status(404).json({ message: 'Ticket not found' });
        }

        res.json({ message: 'Ticket update sucessfully' });
    } catch (error) {
        console.error("Error to update ticket:", error);
        res.status(500).json({ message: 'Error to update ticket' });
    }
};

module.exports = { updateTicket };