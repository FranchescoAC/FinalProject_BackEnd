const connection = require('../config/db'); // Asegúrate de la ruta correcta

// Funciones CRUD usando async/await

const Bus = {
    // Registrar un bus
  updateAvailability: async (id, availability) => {
        try {
            const query = 'UPDATE buses SET availability = ? WHERE id = ?';
            const [result] = await connection.query(query, [availability, id]); // Uso de await
            return result;
        } catch (error) {
            console.error('Error to update bus availability :', error);
            throw error;
        }
    }
};

module.exports = Bus;