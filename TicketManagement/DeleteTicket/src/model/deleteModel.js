const db = require('../config/db'); // Asegúrate de que la ruta sea correcta

const Ticket = (id, bus_id, seat_number, status) => ({
    id,
    bus_id,
    seat_number,
    status,
});

const TicketModel = {
    async delete(id) {
        const connection = await db.getConnection();
        try {
            const [result] = await connection.execute('DELETE FROM tickets WHERE id = ?', [id]);
            return result.affectedRows > 0; // Devuelve true si se eliminó al menos una fila
        } catch (error) {
            console.error('Error en TicketModel.delete:', error);
            throw error; // Re-lanza el error
        } finally {
            connection.release();
        }
    }
}

module.exports = TicketModel;