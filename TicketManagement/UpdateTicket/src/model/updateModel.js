const db = require('../config/db');

const TicketModel = {
    async update(id, updates) {
        const connection = await db.getConnection();
        try {
            const { availabiliti, seat_number } = updates; // Nombre corregido

            const [result] = await connection.execute(
                'UPDATE tickets SET availabiliti = ?, seat_number = ? WHERE id = ?', // Nombre corregido
                [availabiliti, seat_number, id] // Nombre corregido
            );
            return result.affectedRows > 0;
        } catch (error) {
            // ... (resto del código igual)
        } finally {
            // ...
        }
    }
}

module.exports = TicketModel;