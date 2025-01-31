const mongoose = require("mongoose");

const routeSchema = new mongoose.Schema({
  origin: { type: String, required: true },
  destination: { type: String, required: true },
  price: { type: Number, required: true },
  availableSeats: { type: Number, required: true },
  departureTime: { type: Date, required: true },
});

module.exports = mongoose.model("Route", routeSchema);

