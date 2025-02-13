const express = require('express');
const router = express.Router();
const { updateTicket } = require('../controller/updateController');

router.put('/update/:id', updateTicket);

module.exports = router;