const Bus = require('../model/updateModel');
const connection = require('../config/db'); // Ajusta la ruta según corresponda

// Registrar un bus
const updateBusAvailability = async (req, res) => {
    try {
        const { id } = req.params;
        const { availability } = req.body;
        const result = await Bus.updateAvailability(id, availability);
        res.status(200).json({ message: 'Bus availability updated successfully', result });
    } catch (error) {
        console.error('Error updating bus availability:', error);
        res.status(500).json({ error: 'Error updating bus availability' });
    }
};
// Exportar las funciones
module.exports = { updateBusAvailability};
