const express = require('express');
const router = express.Router();
const deleteController = require('../controller/deleteController'); // Asegúrate de la ruta correcta

// Rutas
router.delete('/delete/:id', deleteController.deleteBus);


module.exports = router;
