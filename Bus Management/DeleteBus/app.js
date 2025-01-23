const express = require('express');
const bodyParser = require('body-parser');
const deleteRoutes = require('../DeleteBus/src/routes/deleteRoutes');
const dotenv = require('dotenv');

// Configuración
dotenv.config();
const app = express();
const PORT = 3000;

// Middleware
app.use(bodyParser.json());
app.use('/api/buses', deleteRoutes);

// Iniciar el servidor
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
