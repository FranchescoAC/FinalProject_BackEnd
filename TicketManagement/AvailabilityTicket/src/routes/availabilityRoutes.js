const express = require('express');

const { viewTicketsAvailability } = require('../controller/availabilityController');
router = express.Router();

// Define routes
router.get('/availability', viewTicketsAvailability);


module.exports = router;