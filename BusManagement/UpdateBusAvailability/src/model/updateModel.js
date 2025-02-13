const connection = require("../config/db"); // Asegúrate de que la ruta sea correcta

const Bus = {
    // ✅ Función para actualizar la disponibilidad de un bus
    updateAvailability: async (id, availability) => {
        try {
            const query = "UPDATE buses SET availability = ? WHERE id = ?";
            
            // 🔹 Ejecutar la consulta y obtener los resultados
            const [result] = await connection.promise().query(query, [parseInt(availability), id]); // Convertir a entero

            // 🔹 Verificar si se actualizó alguna fila
            if (result.affectedRows > 0) {
                console.log(`✅ Disponibilidad actualizada para el bus con ID: ${id}`);
                return result;
            } else {
                console.warn(`⚠ No se encontró ningún bus con ID: ${id}`);
                return { affectedRows: 0 };
            }
        } catch (error) {
            console.error("❌ Error al actualizar la disponibilidad del bus:", error);
            throw error;
        }
    }
};

module.exports = Bus;
