const express = require("express");
const { updateRoute } = require("../controller/updateController");

const router = express.Router();

// Ruta para actualizar una ruta existente
router.put("/update/:id", updateRoute);

module.exports = router;
