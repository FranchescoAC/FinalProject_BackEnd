const connection = require('../config/db'); // Asegúrate de la ruta correcta

// Funciones CRUD usando async/await

const Bus = {
    // Registrar un bus
 getAllBuses: async () => {
        try {
            const query = 'SELECT * FROM buses';
            const [rows] = await connection.query(query); // Uso de await
            return rows;
        } catch (error) {
            console.error('Error to get buses:', error);
            throw error;
        }
    }
};

module.exports = Bus;