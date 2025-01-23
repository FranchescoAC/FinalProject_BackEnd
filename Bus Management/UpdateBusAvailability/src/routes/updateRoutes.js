const express = require('express');
const router = express.Router();
const updateController = require('../controller/updateController'); // Asegúrate de la ruta correcta

// Rutas
router.put('/update/:id', updateController.updateBusAvailability);


module.exports = router;
