import axios from 'axios';


export const registerUser = async (userData) => {
  return await axios.post(`http://localhost:2000/register`, userData);
};

export const loginUser = async (credentials) => {
  return await axios.post(`http://localhost:2001/login`, credentials);
};
