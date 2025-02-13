const Bus = require("../model/updateModel");
const connection = require("../config/db"); // Asegúrate de que la ruta sea correcta

// ✅ Actualizar disponibilidad del bus
const updateBusAvailability = async (req, res) => {
    try {
        const { id } = req.params;
        const { availability } = req.body;

        // 🔹 Verificar si availability es un número válido
        if (availability === undefined || availability < 0) {
            return res.status(400).json({ error: "Invalid availability value. Must be 0 or greater." });
        }

        // 🔹 Ejecutar la actualización en la base de datos
        const result = await Bus.updateAvailability(id, availability);

        if (result.affectedRows > 0) {
            res.status(200).json({ message: "Bus availability updated successfully", availability });
        } else {
            res.status(404).json({ error: "Bus not found." });
        }
    } catch (error) {
        console.error("Error updating bus availability:", error);
        res.status(500).json({ error: "Internal server error while updating bus availability." });
    }
};

// ✅ Exportar las funciones
module.exports = { updateBusAvailability };
