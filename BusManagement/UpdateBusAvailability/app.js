const express = require('express');
const bodyParser = require('body-parser');
const updateRoutes = require('../UpdateBusAvailability/src/routes/updateRoutes');
const dotenv = require('dotenv');

// Configuración
dotenv.config();
const app = express();
const PORT = 3002;

// Middleware
app.use(bodyParser.json());

app.use('/api/buses', updateRoutes);


// Iniciar el servidor
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
