import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { AuthProvider, useAuth } from "./context/AuthContext";
import { useContext } from "react";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import Dashboard from "./pages/user/Dashboard";
import Tickets from "./pages/user/Tickets";
import Payments from "./pages/user/Payments";
import DashboardAdmin from "./pages/admin/DashboardAdmin";
import ManageRoutes from "./pages/admin/ManageRoutes";
import ListRoutes from "./pages/admin/ListRoutes";
import UpdateRoute from "./pages/admin/UpdateRoute";
import DeleteRoute from "./pages/admin/DeleteRoute";


function App() {
  const { token, role } = useAuth(); // ✅ Obtenemos token y rol del contexto

  return (
    <BrowserRouter>
      <Routes>
        {/* Rutas de autenticación */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Rutas protegidas para usuarios normales */}
        <Route path="/dashboard" element={token && role === "user" ? <Dashboard /> : <Navigate to="/login" />} />
        <Route path="/tickets" element={token && role === "user" ? <Tickets /> : <Navigate to="/login" />} />

        {/* Rutas protegidas para admin */}
        <Route path="/admin/dashboard" element={token && role === "admin" ? <DashboardAdmin /> : <Navigate to="/login" />} />
        <Route path="/admin/routes" element={token && role === "admin" ? <ManageRoutes /> : <Navigate to="/login" />} />
        <Route path="/admin/routes/list" element={token && role === "admin" ? <ListRoutes /> : <Navigate to="/login" />} />
        <Route path="/admin/routes/update" element={token && role === "admin" ? <UpdateRoute /> : <Navigate to="/login" />} />
        <Route path="/admin/routes/delete" element={token && role === "admin" ? <DeleteRoute /> : <Navigate to="/login" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
