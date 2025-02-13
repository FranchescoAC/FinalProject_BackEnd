require('dotenv').config(); // Carga las variables de entorno (¡primero!)
const express = require("express");
const cors = require("cors");
const connectDB = require("./config/database"); // Importa la conexión a la base de datos
const registerRoute = require("./routes/registerRoute");

const app = express();
const PORT = process.env.PORT || 6001;

// Middleware
app.use(cors());
app.use(express.json());

// Conexión a la base de datos (¡después del middleware!)
connectDB();

// Rutas
app.use("/api/routes", registerRoute);

// Iniciar servidor (¡al final!)
app.listen(PORT, '0.0.0.0', () => {
  console.log(`ListRoute Service running on port ${PORT}`);
});