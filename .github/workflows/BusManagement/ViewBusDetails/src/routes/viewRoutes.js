const express = require('express');
const router = express.Router();
const viewController = require('../controller/viewController'); // Asegúrate de la ruta correcta

// Rutas
router.get('/details', viewController.viewBusDetails);


module.exports = router;
