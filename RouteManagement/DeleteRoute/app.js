require('dotenv').config();
const express = require("express");
const cors = require("cors"); // Importar cors
const connectDB = require("./config/database");
const deleteRoutes = require("./routes/deleteRoutes");

const app = express();
const PORT = process.env.PORT || 6004;

// Middleware
app.use(express.json());
app.use(cors()); // Habilitar CORS

// Database Connection
connectDB();

// Routes
app.use("/api/routes", deleteRoutes);

// Start Server
app.listen(PORT, '0.0.0.0', () => { // Escuchar en 0.0.0.0
  console.log(`DeleteRoute Service running on port ${PORT}`);
});
