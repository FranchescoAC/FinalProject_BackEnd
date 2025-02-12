import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../../services/authService";
import { Container, Card, Form, Button } from "react-bootstrap";

const Register = () => {
  const [credentials, setCredentials] = useState({ username: "", email: "", password: "" });
  const navigate = useNavigate();

  const handleChange = (e) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await registerUser(credentials);
      navigate("/login");
    } catch (error) {
      console.error("Registration failed:", error);
    }
  };

  return (
    <div className="register-container">
      <Container className="d-flex justify-content-center align-items-center vh-100">
        <Card className="p-4 shadow-lg" style={{ width: "400px" }}>
          <h2 className="text-center mb-4">Create an Account</h2>
          <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3">
              <Form.Control type="text" name="username" placeholder="Username" onChange={handleChange} required />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control type="email" name="email" placeholder="Email" onChange={handleChange} required />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control type="password" name="password" placeholder="Password" onChange={handleChange} required />
            </Form.Group>
            <Button variant="primary" type="submit" className="w-100">
              Register
            </Button>
          </Form>
          <p className="text-center mt-3">
            Already have an account? <a href="/login">Login</a>
          </p>
        </Card>
      </Container>
    </div>
  );
};

export default Register;
