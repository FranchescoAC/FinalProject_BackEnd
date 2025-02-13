import axios from "axios";

const API_URL = "http://localhost:2001/api/users"; // Ajusta el puerto si es necesario

// Obtener todos los usuarios
export const getUsers = async () => {
  return await axios.get("http://localhost:2002/users");
};

/* Eliminar un usuario
export const deleteUser = async (userId) => {
    return await axios.delete(`${"http://localhost:2004/deleteUser"}?id=${userId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
  };*/


/* ✅ Actualizar usuario
export const updateUser = async (userId, updatedData) => {
  return await axios.put(`${"http://localhost:2003/deleteUser"}/${userId}`, updatedData);
}; 
*/