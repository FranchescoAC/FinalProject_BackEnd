const express = require('express');
const router = express.Router();
const registerController = require('../controller/registerController'); // Asegúrate de la ruta correcta

// Rutas
router.post('/register', registerController.registerBus);


module.exports = router;
