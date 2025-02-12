const express = require("express");
const router = express.Router();
const updateController = require("../controller/updateController"); // Asegúrate de que la ruta sea correcta

// ✅ Ruta para actualizar la disponibilidad de un bus
router.put("/buses/:id/availability", updateController.updateBusAvailability);

module.exports = router;

