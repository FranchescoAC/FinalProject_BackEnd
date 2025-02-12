import { useAuth } from "../context/AuthContext";

// Si el usuario no está autenticado, redirigir al login
if (!token) {
    navigate("/login");
    return null;
  }