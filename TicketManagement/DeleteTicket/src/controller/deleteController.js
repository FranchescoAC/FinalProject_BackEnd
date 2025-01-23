const TicketModel = require('../model/deleteModel');

const deleteTicket = async (req, res) => {
    try {
        const { id } = req.params;
        const deleted = await TicketModel.delete(id);
        if (!deleted) {
            return res.status(404).json({ message: 'Ticket not found' });
        }
        res.json({ message: 'Ticket delete sucessfully' });
    } catch (error) {
        console.error("Error to delete ticket:", error);
        res.status(500).json({ message: 'Error to delete ticket' });
    }
};


module.exports = { deleteTicket };