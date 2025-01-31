const express = require("express");
const { listRoutes } = require("../controller/listController");

const router = express.Router();

router.get("/list", listRoutes);

module.exports = router;
