import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./ManageRoutes.css";

function UpdateRoute() {
  const [routeId, setRouteId] = useState("");
  const [updatedRoute, setUpdatedRoute] = useState({
    origin: "",
    destination: "",
    price: "",
    availableSeats: "",
    departureTime: ""
  });
  const navigate = useNavigate();

  const handleUpdate = async () => {
    try {
      await axios.put(`http://98.82.125.215:6003/api/routes/update/${routeId}`, updatedRoute);
      alert("Route updated successfully");
      navigate("/admin/routes");
    } catch (error) {
      console.error("Error updating route", error);
    }
  };

  return (
    <div className="manage-routes-container">
      <h1>Update Route</h1>
      <input value={routeId} onChange={(e) => setRouteId(e.target.value)} placeholder="Route ID" />
      <input value={updatedRoute.origin} onChange={(e) => setUpdatedRoute({ ...updatedRoute, origin: e.target.value })} placeholder="Origin" />
      <input value={updatedRoute.destination} onChange={(e) => setUpdatedRoute({ ...updatedRoute, destination: e.target.value })} placeholder="Destination" />
      <input value={updatedRoute.price} type="number" onChange={(e) => setUpdatedRoute({ ...updatedRoute, price: e.target.value })} placeholder="Price" />
      <input value={updatedRoute.availableSeats} type="number" onChange={(e) => setUpdatedRoute({ ...updatedRoute, availableSeats: e.target.value })} placeholder="Available Seats" />
      <input value={updatedRoute.departureTime} type="datetime-local" onChange={(e) => setUpdatedRoute({ ...updatedRoute, departureTime: e.target.value })} placeholder="Departure Time" />
      <button className="update-route-button" onClick={handleUpdate}>Update Route</button>
      <button className="back-button" onClick={() => navigate("/admin/routes")}>Back to Manage Routes</button>
    </div>
  );
}

export default UpdateRoute;
