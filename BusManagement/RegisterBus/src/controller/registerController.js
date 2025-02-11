const Bus = require('../model/registerModel');
const connection = require('../config/db');

const registerBus = async (req, res) => {
    try {
        const { name, capacity } = req.body;

        // 1. Log detallado de los datos recibidos
        console.log("Datos recibidos:", { name, capacity });

        const result = await Bus.registerBus(name, capacity);

        // 2. Log del resultado de la operación (si es necesario)
        console.log("Resultado de la operación:", result);

        res.status(201).json({ message: 'Bus registered successfully', result });
    } catch (error) {
        // 3. Log del error *completo* con stack trace
        console.error('Error registering bus:', error);
        console.error('Stack trace:', error.stack); // Información adicional del error

        // 4. Envía un mensaje de error *más específico* al cliente (opcional)
        res.status(500).json({ error: 'Error registering bus. Por favor, inténtelo nuevamente más tarde.' }); // Mensaje genérico para no exponer detalles internos
    }
};

module.exports = { registerBus };