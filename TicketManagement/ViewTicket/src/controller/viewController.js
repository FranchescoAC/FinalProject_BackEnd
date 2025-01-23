const TicketModel = require('../model/viewModel');

const getAllTickets = async (req, res) => {
    try {
        const allTickets = await TicketModel.getAll();
        res.status(200).json(allTickets);
    } catch (error) {
        console.error("Error to get tickets:", error);
        res.status(500).json({ message: 'Error to get tickets' });
    }
};

module.exports = { getAllTickets };