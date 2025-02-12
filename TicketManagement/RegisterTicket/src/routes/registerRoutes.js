const express = require('express');
const router = express.Router();
const { registerTicket } = require('../controller/registerController');

router.post('/register', registerTicket);

module.exports = router;