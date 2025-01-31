const Route = require("../model/updateModel");

const updateRoute = async (req, res) => {
  try {
    const { id } = req.params; // ID del documento a actualizar
    const updatedData = req.body; // Nuevos datos para actualizar

    // Busca el documento por ID y actualiza los campos con los datos nuevos
    const updatedRoute = await Route.findByIdAndUpdate(id, updatedData, { new: true });

    if (!updatedRoute) {
      return res.status(404).json({ message: "Route not found" });
    }

    res.status(200).json({ message: "Route updated successfully", route: updatedRoute });
  } catch (err) {
    res.status(500).json({ message: "Error updating route", error: err.message });
  }
};

module.exports = { updateRoute };

