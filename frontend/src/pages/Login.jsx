import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { loginUser } from "../services/authService";
import { useAuth } from "../context/AuthContext";
import { Form, Button, Container, Card } from "react-bootstrap";

const Login = () => {
  const [credentials, setCredentials] = useState({ email: "", password: "" });
  const { login, token } = useAuth(); // Obtiene el token del contexto
  const navigate = useNavigate();

  // ✅ Si ya hay un token, redirigir dentro de useEffect()
  useEffect(() => {
    if (token) {
      navigate("/dashboard");
    }
  }, [token, navigate]);

  const handleChange = (e) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await loginUser(credentials);
      login(response.data.token);
      navigate("/dashboard"); // Redirigir después de iniciar sesión exitosamente
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
