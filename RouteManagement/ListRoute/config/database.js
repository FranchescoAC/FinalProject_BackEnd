const mongoose = require("mongoose");

const connectDB = async () => {
  try {
    const dbURI = process.env.MONGO_URI || "mongodb://localhost:27017/RouteManagement";
    await mongoose.connect(dbURI, { useNewUrlParser: true, useUnifiedTopology: true });
    console.log("MongoDB connected for ListRoute");
  } catch (err) {
    console.error("Database connection error:", err.message);
    process.exit(1);
  }
};

module.exports = connectDB;

