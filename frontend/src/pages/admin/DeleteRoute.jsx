import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./ManageRoutes.css";

function DeleteRoute() {
  const [routeId, setRouteId] = useState("");
  const navigate = useNavigate();

  const handleDelete = async () => {
    try {
      await axios.delete(`http://44.196.89.150:6004/api/routes/delete/${routeId}`);
      alert("Route deleted successfully");
      navigate("/admin/routes");
    } catch (error) {
      console.error("Error deleting route", error);
    }
  };

  return (
    <div className="manage-routes-container">
      <h1>Delete Route</h1>
      <input value={routeId} onChange={(e) => setRouteId(e.target.value)} placeholder="Route ID" />
      <button className="delete-route-button" onClick={handleDelete}>Delete Route</button>
      <button className="back-button" onClick={() => navigate("/admin/routes")}>Back to Manage Routes</button>
    </div>
  );
}

export default DeleteRoute;
