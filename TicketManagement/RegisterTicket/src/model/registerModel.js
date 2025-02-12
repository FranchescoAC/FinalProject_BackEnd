const db = require('../config/db');

const TicketModel = {
    async create(ticketData) {
        const connection = await db.getConnection();
        try {
            const { name_bus, seat_number, availabiliti } = ticketData; // Nombre corregido

            const [result] = await connection.execute(
                'INSERT INTO tickets (name_bus, seat_number, availabiliti) VALUES (?, ?, ?)', // Nombre corregido
                [name_bus, seat_number, availabiliti] // Nombre corregido
            );
            return { id: result.insertId, name_bus, seat_number, availabiliti }; // Nombre corregido
        } catch (error) {
            // ... (resto del código igual)
        } finally {
            // ...
        }
    }
};

module.exports = TicketModel;