import { useNavigate } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";

function DashboardAdmin() {
  const navigate = useNavigate();
  const { logout } = useAuth();

  return (
    <div>
      <h1>Admin Dashboard</h1>
      <button onClick={() => navigate("/admin/routes")}>Manage Routes</button>
      <button onClick={logout}>Logout</button>
    </div>
  );
}

export default DashboardAdmin;
