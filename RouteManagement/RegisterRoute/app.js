const express = require("express");
const connectDB = require("./config/database");
const registerRoute = require("./routes/registerRoute");

const app = express();
const PORT = process.env.PORT || 3001;

// Middleware
app.use(express.json());

// Database Connection
connectDB();

// Routes
app.use("/api/routes", registerRoute);

// Start Server
app.listen(PORT, () => {
  console.log(`RegisterRoute Service running on port ${PORT}`);
});
