const Bus = require('../model/registerModel');
const connection = require('../config/db'); // Ajusta la ruta según corresponda

// Registrar un bus
const registerBus = async (req, res) => {
    try {
        const { name, capacity } = req.body;
        const result = await Bus.registerBus(name, capacity);
        res.status(201).json({ message: 'Bus registered successfully', result });
    } catch (error) {
        console.error('Error registering bus:', error);
        res.status(500).json({ error: 'Error registering bus' });
    }
};
// Exportar las funciones
module.exports = { registerBus};
