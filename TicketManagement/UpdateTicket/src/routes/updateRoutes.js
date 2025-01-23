const express = require('express');
const { updateTicket } = require('../controller/updateController');

router = express.Router();

// Define routes
router.put('/update/:id', updateTicket);                  // Update a ticket


module.exports = router;