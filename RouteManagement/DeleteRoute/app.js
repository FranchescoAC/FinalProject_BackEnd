const express = require("express");
const connectDB = require("./config/database");
const deleteRoutes = require("./routes/deleteRoutes");

const app = express();
const PORT = process.env.PORT || 3004;

// Middleware
app.use(express.json());

// Database Connection
connectDB();

// Routes
app.use("/api/routes", deleteRoutes);

// Start Server
app.listen(PORT, () => {
  console.log(`DeleteRoute Service running on port ${PORT}`);
});

