const connection = require('../config/db'); // Asegúrate de la ruta correcta

// Funciones CRUD usando async/await

const Bus = {
    // Registrar un bus
    registerBus: async (name, capacity) => {
        try {
            const query = 'INSERT INTO buses (name, capacity) VALUES (?, ?)';
            const [result] = await connection.query(query, [name, capacity]); // Uso de await
            return result;
        } catch (error) {
            console.error('Error to register bus:', error);
            throw error;
        }
    }
};

module.exports = Bus;