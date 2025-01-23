const express = require('express');

const { registerTicket } = require('../controller/registerController');

// Define routes
router.post('/register', registerTicket);          // Register a new ticketS

module.exports = router;