import axios from 'axios';

const API_URL = 'http://localhost:4002'; // TicketManagement

export const getRoutes = async () => {
    const response = await axios.get(`http://localhost:3002/api/routes/list`);
};

export const getBuses = async () => {
  return await axios.get(`${API_URL}/buses`);
};

export const bookTicket = async (ticketData) => {
  return await axios.post(`${API_URL}/registerTicket`, ticketData);
};
