import axios from 'axios';

const API_URL = 'http://localhost:4002'; // TicketManagement

export const getRoutes = async () => {
  try {
    const response = await axios.get('http://localhost:6002/api/routes/list');
    return response.data.routes; // Devuelve solo el array de rutas
  } catch (error) {
    console.error('Error fetching routes:', error);
    return []; // Devuelve un array vacío en caso de error
  }
};

export const getBuses = async () => {
  try {
    const response = await axios.get(`http://3.224.72.234:3003/api/buses/details`);
    return response.data; // Devuelve los datos de los buses
  } catch (error) {
    console.error('Error fetching buses:', error);
    return []; // Devuelve un array vacío en caso de error
  }
};

export const bookTicket = async (ticketData) => {
  return await axios.post(`${API_URL}/registerTicket`, ticketData);
};
