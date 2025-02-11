import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import "../styles/Dashboard.css";

const Dashboard = () => {
  const { token, logout } = useAuth();
  const navigate = useNavigate();

  // Si el usuario no está autenticado, redirigir al login
  if (!token) {
    navigate("/login");
    return null;
  }

  return (
    <div className="dashboard">
      {/* Sidebar de navegación */}
      <aside className="sidebar">
        <h2>Dashboard</h2>
        <ul>
          <li onClick={() => navigate("/tickets")}>🎟 Tickets</li>
          <li onClick={() => navigate("/notifications")}>🔔 Notifications</li>
          <li onClick={() => navigate("/payments")}>💳 Payments</li>
        </ul>
        <button className="logout-btn" onClick={logout}>🚪 Logout</button>
      </aside>

      {/* Contenido principal */}
      <main className="main-content">
        <h1>Welcome to Your Dashboard</h1>
        <p>Manage your tickets, payments, and notifications here.</p>
      </main>
    </div>
  );
};

export default Dashboard;
