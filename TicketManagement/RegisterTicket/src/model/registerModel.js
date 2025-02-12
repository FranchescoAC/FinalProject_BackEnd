const db = require('../config/db');

const TicketModel = {
    async create(ticketData) {
        const connection = await db.getConnection();
        try {
            const { name_bus, seat_number } = ticketData;

            const [result] = await connection.execute(
                'INSERT INTO tickets (name_bus, seat_number) VALUES (?, ?)',
                [name_bus, seat_number]
            );
            return { id: result.insertId, name_bus, seat_number };
        } catch (error) {
            console.error('❌ Error en TicketModel.create:', error);
            throw error;
        } finally {
            connection.release();
        }
    }
};

module.exports = TicketModel;