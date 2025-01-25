const express = require("express");
const connectDB = require("./config/database");
const listRoutes = require("./routes/listRoutes");

const app = express();
const PORT = process.env.PORT || 3002;

// Middleware
app.use(express.json());

// Database Connection
connectDB();

// Routes
app.use("/api/routes", listRoutes);

// Start Server
app.listen(PORT, () => {
  console.log(`ListRoute Service running on port ${PORT}`);
});

