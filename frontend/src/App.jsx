import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { AuthProvider, useAuth } from "./context/AuthContext";
import { useContext } from "react";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import Dashboard from "./pages/user/Dashboard";
import Tickets from "./pages/user/Tickets";
import Payments from "./pages/user/Payments";
import DashboardAdmin from "./pages/admin/DashboardAdmin";

function App() {
  const { token, role } = useAuth(); // ✅ Ahora obtenemos el rol del contexto

  return (
    <BrowserRouter>
      <Routes>
        {/* Rutas de autenticación */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Rutas protegidas para usuarios normales */}
        <Route path="/dashboard" element={token && role === "user" ? <Dashboard /> : <Navigate to="/login" />} />
        <Route path="/tickets" element={token && role === "user" ? <Tickets /> : <Navigate to="/login" />} />
        <Route path="/payments" element={token && role === "user" ? <Payments /> : <Navigate to="/login" />} />

        {/* Rutas protegidas para administradores */}
        <Route path="/admin/dashboard" element={token && role === "admin" ? <DashboardAdmin /> : <Navigate to="/login" />} />

        {/* Redirección por defecto */}
        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
