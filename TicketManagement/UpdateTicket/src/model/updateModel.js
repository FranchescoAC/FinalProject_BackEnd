const db = require('../config/db'); // Asegúrate de que la ruta sea correcta

const Ticket = (id, bus_id, seat_number, status) => ({
    id,
    bus_id,
    seat_number,
    status,
});

const TicketModel = {
    async update(id, updates) {
        const connection = await db.getConnection();
        try {
            const [result] = await connection.execute(
                'UPDATE tickets SET status = ?, seat_number = ? WHERE id = ?',
                [updates.status, updates.seat_number, id]
            );
            return result.affectedRows > 0; // Devuelve true si se actualizó al menos una fila
        } catch (error) {
            console.error('Error en TicketModel.update:', error);
            throw error; // Re-lanza el error
        } finally {
            connection.release();
        }
    }
}
module.exports = TicketModel;