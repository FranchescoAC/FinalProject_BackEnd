require('dotenv').config();
const express = require("express");
const connectDB = require("./config/database");
const updateRoutes = require("./routes/updateRoutes");

const app = express();
const PORT = process.env.PORT || 6003;

// Middleware
app.use(express.json());

// Database Connection
connectDB();

// Routes
app.use("/api/routes", updateRoutes);

// Start Server
app.listen(PORT, '0.0.0.0', () => { // Escuchar en 0.0.0.0
  console.log(`ListRoute Service running on port ${PORT}`);
});



