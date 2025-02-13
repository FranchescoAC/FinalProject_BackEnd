const mongoose = require("mongoose");
require("dotenv").config(); // Cargar variables de entorno

const connectDB = async () => {
  try {
    const dbURI = process.env.MONGO_URI;

    if (!dbURI) {
      throw new Error("MONGO_URI no está definido en el archivo .env");
    }

    await mongoose.connect(dbURI, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
      serverSelectionTimeoutMS: 60000,
      socketTimeoutMS: 45000,
      connectTimeoutMS: 30000,
    });

    console.log("✅ Conectado a MongoDB");
  } catch (err) {
    console.error("❌ Error de conexión a MongoDB:", err.message);
    process.exit(1); // Cerrar la app si hay error
  }
};

module.exports = connectDB;
