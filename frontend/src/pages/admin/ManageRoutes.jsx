import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./ManageRoutes.css";

function ManageRoutes() {
  const [routes, setRoutes] = useState([]);
  const [newRoute, setNewRoute] = useState({
    origin: "",
    destination: "",
    price: "",
    availableSeats: "",
    departureTime: ""
  });

  const navigate = useNavigate();

  useEffect(() => {
    fetchRoutes();
  }, []);

  const fetchRoutes = async () => {
    try {
      const response = await axios.get("http://98.82.125.215:/api/routes");
      setRoutes(response.data);
    } catch (error) {
      console.error("Error fetching routes", error);
    }
  };

  const addRoute = async () => {
    try {
      await axios.post("http://44.196.89.150:6001/api/routes/register", newRoute);
      fetchRoutes();
      setNewRoute({ origin: "", destination: "", price: "", availableSeats: "", departureTime: "" });
    } catch (error) {
      console.error("Error adding route", error);
    }
  };

  return (
    <div className="manage-routes-container">
      <h1>Manage Routes</h1>
      <div className="form-container">
        <input value={newRoute.origin} onChange={(e) => setNewRoute({ ...newRoute, origin: e.target.value })} placeholder="Origin" />
        <input value={newRoute.destination} onChange={(e) => setNewRoute({ ...newRoute, destination: e.target.value })} placeholder="Destination" />
        <input value={newRoute.price} type="number" onChange={(e) => setNewRoute({ ...newRoute, price: e.target.value })} placeholder="Price" />
        <input value={newRoute.availableSeats} type="number" onChange={(e) => setNewRoute({ ...newRoute, availableSeats: e.target.value })} placeholder="Available Seats" />
        <input value={newRoute.departureTime} type="datetime-local" onChange={(e) => setNewRoute({ ...newRoute, departureTime: e.target.value })} placeholder="Departure Time" />
        <button className="add-route-button" onClick={addRoute}>Add Route</button>
      </div>
      <ul className="routes-list">
        {routes.map((route) => (
          <li key={route.id} className="route-item">
            <span>{route.origin} to {route.destination} - ${route.price} - {route.availableSeats} seats - {new Date(route.departureTime).toLocaleString()}</span>
          </li>
        ))}
      </ul>
      <div className="routes-actions">
        <button className="action-button" onClick={() => navigate("/admin/routes/list")}>View Routes</button>
        <button className="action-button" onClick={() => navigate("/admin/routes/update")}>Update Route</button>
        <button className="action-button" onClick={() => navigate("/admin/routes/delete")}>Delete Route</button>
      </div>
      <button className="back-button" onClick={() => navigate("/admin/dashboard")}>Back to Dashboard</button>
    </div>
  );
}

export default ManageRoutes;
