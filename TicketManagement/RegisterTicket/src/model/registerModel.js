const db = require('../config/db'); // Asegúrate de que la ruta sea correcta

const Ticket = (id, user_id, bus_id, seat_number, status) => ({
    id,
    user_id,
    bus_id,
    seat_number,
    status,
});

const TicketModel = {
    async create(ticketData) {
        const connection = await db.getConnection();
        try {
            const [result] = await connection.execute(
                'INSERT INTO tickets (user_id, bus_id, seat_number, status) VALUES (?, ?, ?, ?)',
                [ticketData.user_id, ticketData.bus_id, ticketData.seat_number, ticketData.status]
            );
            return { id: result.insertId, ...ticketData }; // Devuelve un objeto con el ID
        } catch (error) {
            console.error('❌ Error en TicketModel.create:', error);
            throw error; // Re-lanza el error
        } finally {
            connection.release();
        }
    }
};

module.exports = TicketModel;
