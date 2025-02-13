import React from "react";
import { Container, Row, Col, Card, Button } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

const DashboardAdmin = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token"); // 🔥 Eliminar el token JWT
    localStorage.removeItem("role"); // 🔥 Eliminar el rol del usuario
    navigate("/login"); // 🔥 Redirigir al login
  };

  return (
    <Container className="mt-5">
      <h2 className="text-center mb-4">Admin Dashboard</h2>

      <Row className="justify-content-center">
        <Col md={6}>
          <Card className="shadow mb-4">
            <Card.Body>
              <h5 className="text-center">Manage Users</h5>
              <Button variant="primary" onClick={() => navigate("/admin/manage-users")} block>
                Go to Users
              </Button>
            </Card.Body>
          </Card>
        </Col>

        <Col md={6}>
          <Card className="shadow mb-4">
            <Card.Body>
              <h5 className="text-center">Manage Buses</h5>
              <Button variant="primary" onClick={() => navigate("/admin/manage-buses")} block>
                Go to Buses
              </Button>
            </Card.Body>
          </Card>
        </Col>

        <Col md={6}>
          <Card className="shadow mb-4">
            <Card.Body>
              <h5 className="text-center">Manage Tickets</h5>
              <Button variant="primary" onClick={() => navigate("/admin/manage-tickets")} block>
                Go to Tickets
              </Button>
            </Card.Body>
          </Card>
        </Col>

        <Col md={6}>
          <Card className="shadow mb-4">
            <Card.Body>
              <h5 className="text-center">Manage Payments</h5>
              <Button variant="primary" onClick={() => navigate("/admin/manage-payments")} block>
                Go to Payments
              </Button>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* 🔥 Botón de Logout */}
      <div className="text-center mt-4">
        <Button variant="danger" onClick={handleLogout}>
          Logout
        </Button>
      </div>
    </Container>
  );
};

export default DashboardAdmin;
