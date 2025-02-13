import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./ManageRoutes.css";

function ListRoutes() {
  const [routes, setRoutes] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchRoutes();
  }, []);

  const fetchRoutes = async () => {
    try {
      const response = await axios.get("http://98.82.125.215:6002/api/routes/list");

      // Acceder correctamente al array de rutas dentro de "routes"
      if (Array.isArray(response.data.routes)) {
        setRoutes(response.data.routes);
      } else {
        console.error("Unexpected response format", response.data);
        setRoutes([]); // Evita errores asignando un array vacío en caso de estructura inesperada
      }
    } catch (error) {
      console.error("Error fetching routes", error);
      setRoutes([]); // Manejo de error: asignamos un array vacío
    }
  };

  return (
    <div className="manage-routes-container">
      <h1>List of Routes</h1>
      <ul className="routes-list">
        {routes.map((route) => (
          <li key={route._id} className="route-item">
            <span>{route.origin} to {route.destination} - ${route.price} - {route.availableSeats} seats - {new Date(route.departureTime).toLocaleString()}</span>
          </li>
        ))}
      </ul>
      <button className="back-button" onClick={() => navigate("/admin/routes")}>Back to Manage Routes</button>
    </div>
  );
}

export default ListRoutes;
