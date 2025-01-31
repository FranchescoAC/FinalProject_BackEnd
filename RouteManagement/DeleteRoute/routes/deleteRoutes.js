const express = require("express");
const { deleteRoute } = require("../controller/deleteController");

const router = express.Router();

// Ruta para eliminar una ruta por ID
router.delete("/delete/:id", deleteRoute);

module.exports = router;
