const express = require('express');
const { getAllTickets } = require('../controller/viewController');
router = express.Router();

// Define routes
router.get('/all', getAllTickets); // View tickets availability


module.exports = router;