const express = require('express');
const { deleteTicket } = require('../controller/deleteController');
router = express.Router();

// Define routes
router.delete('/delete/:id', deleteTicket);               // Delete a ticket

module.exports = router;