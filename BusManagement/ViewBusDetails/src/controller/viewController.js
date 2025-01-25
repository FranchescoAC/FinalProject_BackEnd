const Bus = require('../model/viewModel');
const connection = require('../config/db'); // Ajusta la ruta según corresponda

// Registrar un bus
const viewBusDetails = async (req, res) => {
    try {
        const buses = await Bus.getAllBuses();
        res.status(200).json(buses);
    } catch (error) {
        console.error('Error fetching bus details:', error);
        res.status(500).json({ error: 'Error fetching bus details' });
    }
};
// Exportar las funciones
module.exports = { viewBusDetails};
