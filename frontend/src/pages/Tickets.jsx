import { useEffect, useState } from "react";
import { getRoutes } from "../services/ticketService"; // Importa el servicio para obtener rutas
import { Container, Row, Col, Button, Card, Form } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

const Tickets = () => {
  const [routes, setRoutes] = useState([]); // Estado para guardar las rutas
  const [selectedRoute, setSelectedRoute] = useState(""); // Estado para la ruta seleccionada
  const navigate = useNavigate();

  useEffect(() => {
    // Recuperamos las rutas al cargar la vista
    const fetchRoutes = async () => {
      try {
        const routesData = await getRoutes(); // Llamada al servicio
        setRoutes(routesData); // Guardamos las rutas
      } catch (error) {
        console.error("Error fetching routes:", error);
      }
    };
    fetchRoutes();
  }, []);

  const handleRouteChange = (e) => {
    setSelectedRoute(e.target.value); // Actualiza la ruta seleccionada
  };

  const handleSubmit = () => {
    if (selectedRoute) {
      // Navegar a la página de "Payment" con la ruta seleccionada
      navigate("/payments");
    } else {
      alert("Please select a route!");
    }
  };

  return (
    <Container className="mt-5">
      <h2>Select Route and Bus</h2>

      <Row className="mt-4">
        <Col md={6}>
          <Card className="p-4">
            <h3>Select a Route</h3>
            <Form>
              <Form.Group controlId="formRoute">
                <Form.Label>Available Routes</Form.Label>
                <Form.Control
                  as="select"
                  onChange={handleRouteChange}
                  value={selectedRoute}
                >
                  <option value="">Choose a Route</option>
                  {routes.map((route) => (
                    <option key={route._id} value={route._id}>
                      {route.origin} to {route.destination} - Price: ${route.price}
                    </option>
                  ))}
                </Form.Control>
              </Form.Group>

              <Button variant="primary" onClick={handleSubmit} className="mt-3">
                Continue to Payment
              </Button>
            </Form>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Tickets;
