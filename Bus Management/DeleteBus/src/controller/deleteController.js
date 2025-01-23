const Bus = require('../model/deleteModel');
const connection = require('../config/db'); // Ajusta la ruta según corresponda

// Registrar un bus
const deleteBus = async (req, res) => {
    try {
        const { id } = req.params;
        const result = await Bus.deleteBus(id);
        res.status(200).json({ message: 'Bus deleted successfully', result });
    } catch (error) {
        console.error('Error deleting bus:', error);
        res.status(500).json({ error: 'Error deleting bus' });
    }
};
// Exportar las funciones
module.exports = { deleteBus};
