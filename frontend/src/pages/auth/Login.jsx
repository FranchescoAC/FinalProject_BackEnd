import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { loginUser } from "../../services/authService";
import { useAuth } from "../../context/AuthContext";
import { Form, Button, Container, Card } from "react-bootstrap";

const Login = () => {
  const [credentials, setCredentials] = useState({ email: "", password: "" });
  const { login, token, role } = useAuth();
  const navigate = useNavigate();

  // ✅ Evitar el bucle infinito con `setTimeout`
  useEffect(() => {
    if (token && role) {
      setTimeout(() => {
        if (role === "admin") {
          navigate("/admin/dashboard");
        } else {
          navigate("/dashboard");
        }
      }, 100); // Pequeño delay para evitar el bucle
    }
  }, [token, role]);

  const handleChange = (e) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await loginUser(credentials);
      
      console.log("Login Response:", response.data); // 🔹 Verificar datos devueltos por la API
  
      if (response.data.token && response.data.role) {
        login(response.data.token, response.data.role);
        console.log("Token Guardado:", response.data.token);
        console.log("Rol Guardado:", response.data.role);
        
        // ✅ Redirigir según el rol
        if (response.data.role === "admin") {
          navigate("/admin/dashboard");
        } else {
          navigate("/dashboard");
        }
      } else {
        console.error("Error: No se recibió token o role");
      }
    } catch (error) {
      console.error("Login failed:", error);
    }
  };
  

  return (
    <Container className="d-flex justify-content-center align-items-center vh-100">
      <Card className="p-4 shadow-lg" style={{ width: "400px" }}>
        <h2 className="text-center mb-4">Login</h2>
        <Form onSubmit={handleSubmit}>
          <Form.Group className="mb-3">
            <Form.Control type="email" name="email" placeholder="Email" onChange={handleChange} required />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Control type="password" name="password" placeholder="Password" onChange={handleChange} required />
          </Form.Group>
          <Button variant="success" type="submit" className="w-100">
            Login
          </Button>
        </Form>
        <p className="text-center mt-3">
          Don't have an account? <a href="/register">Register</a>
        </p>
      </Card>
    </Container>
  );
};

export default Login;
