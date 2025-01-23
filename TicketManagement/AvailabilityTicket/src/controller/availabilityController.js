const TicketModel = require('../model/availabilityModel');

const viewTicketsAvailability = async (req, res) => {
    try {
        const availableTickets = await TicketModel.getAvailableTickets();
        res.status(200).json(availableTickets);
    } catch (error) {
        console.error("Error al obtener los tickets disponibles:", error);
        res.status(500).json({ message: 'Error al obtener los tickets disponibles' });
    }
};

module.exports = { viewTicketsAvailability };