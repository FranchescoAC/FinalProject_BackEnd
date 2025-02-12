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
app.listen(PORT, () => {
  console.log(`UpdateRoute Service running on port ${PORT}`);
});


