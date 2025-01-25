const connection = require('../config/db'); // Asegúrate de la ruta correcta

// Funciones CRUD usando async/await

const Bus = {
    // Registrar un bus
 deleteBus: async (id) => {
        try {
            const query = 'DELETE FROM buses WHERE id = ?';
            const [result] = await connection.query(query, [id]); // Uso de await
            return result;
        } catch (error) {
            console.error('Error to delete bus:', error);
            throw error;
        }
    }
};

module.exports = Bus;