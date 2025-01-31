const Route = require("../model/registerModel");

const registerRoute = async (req, res) => {
  try {
    const { origin, destination, price, availableSeats, departureTime } = req.body;

    const newRoute = new Route({ origin, destination, price, availableSeats, departureTime });
    await newRoute.save();

    res.status(201).json({ message: "Route registered successfully", route: newRoute });
  } catch (err) {
    res.status(500).json({ message: "Error registering route", error: err.message });
  }
};

module.exports = { registerRoute };
