const Route = require("../model/listModel");

const listRoutes = async (req, res) => {
  try {
    const routes = await Route.find({});
    res.status(200).json({ routes });
  } catch (err) {
    res.status(500).json({ message: "Error retrieving routes", error: err.message });
  }
};

module.exports = { listRoutes };
