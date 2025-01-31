const express = require("express");
const { registerRoute } = require("../controller/registerController");

const router = express.Router();

router.post("/register", registerRoute);

module.exports = router;
