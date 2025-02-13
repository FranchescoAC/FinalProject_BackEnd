const db = require('../config/db');

const TicketModel = {
    async getAvailableTickets() {
        const connection = await db.getConnection();
        try {
            const [rows] = await connection.execute('SELECT * FROM tickets WHERE availabiliti = ?', [true]); // Usa availabiliti = true
            return rows; // Devuelve directamente las filas
        } catch (error) {
            console.error('Error en TicketModel.getAvailableTickets:', error);
            throw error;
        } finally {
            connection.release();
        }
    },
};

module.exports = TicketModel;