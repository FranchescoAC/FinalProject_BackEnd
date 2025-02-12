const express = require("express");
const cors = require("cors"); // Importa el paquete cors
const connectDB = require("./config/database");
const listRoutes = require("./routes/listRoutes");

const app = express();
const PORT = process.env.PORT || 6002;

// Middleware
app.use(cors()); // Habilita CORS para todas las rutas
app.use(express.json());

// Database Connection
connectDB();

// Routes
app.use("/api/routes", listRoutes);

// Start Server
app.listen(PORT, () => {
  console.log(`ListRoute Service running on port ${PORT}`);
});