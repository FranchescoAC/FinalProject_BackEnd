const db = require('../config/db'); // Asegúrate de que la ruta sea correcta

const Ticket = (id, bus_id, seat_number, status) => ({
    id,
    bus_id,
    seat_number,
    status,
});

const TicketModel = {
  async getAvailableTickets() {
        const connection = await db.getConnection();
        try {
            const [rows] = await connection.execute('SELECT * FROM tickets WHERE status = ?', ['available']);
            return rows.map(row => Ticket(row.id, row.bus_id, row.seat_number, row.status));
        } catch (error) {
            console.error('Error en TicketModel.getAvailableTickets:', error);
            throw error; // Re-lanza el error
        } finally {
            connection.release();
        }
    },
};

module.exports = TicketModel;