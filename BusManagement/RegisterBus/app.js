const express = require('express');
const bodyParser = require('body-parser');
const registerRoutes = require('./src/routes/registerRoutes');
const dotenv = require('dotenv');
// Configuración
dotenv.config();
const app = express();
const PORT = 3001;

// Middleware
app.use(bodyParser.json());
app.use('/api/buses', registerRoutes);

// Iniciar el servidor
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
