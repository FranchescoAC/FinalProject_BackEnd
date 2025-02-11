import { createContext, useContext, useState } from "react";

const AuthContext = createContext(); // Asegúrate de que esto está definido correctamente

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(localStorage.getItem("token") || "");

  const login = (newToken) => {
    setToken(newToken);
    localStorage.setItem("token", newToken);
  };

  const logout = () => {
    setToken("");
    localStorage.removeItem("token");
  };

  return (
    <AuthContext.Provider value={{ token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// Esta línea permite acceder al contexto desde cualquier componente
export const useAuth = () => useContext(AuthContext);
