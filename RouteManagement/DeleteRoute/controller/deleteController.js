const Route = require("../model/deleteModel");

const deleteRoute = async (req, res) => {
  try {
    const { id } = req.params; // ID del documento a eliminar

    // Busca el documento por ID y lo elimina
    const deletedRoute = await Route.findByIdAndDelete(id);

    if (!deletedRoute) {
      return res.status(404).json({ message: "Route not found" });
    }

    res.status(200).json({ message: "Route deleted successfully" });
  } catch (err) {
    res.status(500).json({ message: "Error deleting route", error: err.message });
  }
};

module.exports = { deleteRoute };

