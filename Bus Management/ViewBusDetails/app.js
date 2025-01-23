const express = require('express');
const bodyParser = require('body-parser');
const viewRoutes = require('./src/routes/viewRoutes');

const dotenv = require('dotenv');

// Configuración
dotenv.config();
const app = express();
const PORT = 3003;

// Middleware
app.use(bodyParser.json());
app.use('/api/buses', viewRoutes);


// Iniciar el servidor
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
