import { useEffect, useState } from "react";
import { getRoutes, getBuses } from "../services/ticketService"; // Importa los servicios
import { Container, Row, Col, Button, Card, Form, Alert, Spinner } from "react-bootstrap";
import { useNavigate } from "react-router-dom";


const Tickets = () => {
  const [routes, setRoutes] = useState([]); // Estado para las rutas
  const [buses, setBuses] = useState([]); // Estado para los buses
  const [selectedRoute, setSelectedRoute] = useState(""); // Ruta seleccionada
  const [selectedBus, setSelectedBus] = useState(""); // Bus seleccionado
  const [loading, setLoading] = useState(false); // Estado para el spinner de carga
  const [error, setError] = useState(""); // Estado para manejar errores
  const navigate = useNavigate();

  // Si el usuario no está autenticado, redirigir al login

  // Obtener las rutas al cargar la vista
  useEffect(() => {
    const fetchRoutes = async () => {
      setLoading(true);
      try {
        const routesData = await getRoutes(); // Llamada al servicio
        setRoutes(routesData); // Guardamos las rutas
      } catch (error) {
        setError("Error fetching routes. Please try again later.");
        console.error("Error fetching routes:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchRoutes();
  }, []);

  // Obtener los buses cuando se selecciona una ruta
  useEffect(() => {
    if (selectedRoute) {
      const fetchBuses = async () => {
        setLoading(true);
        try {
          const busesData = await getBuses(); // Llamada al servicio
          setBuses(busesData); // Guardamos los buses
        } catch (error) {
          setError("Error fetching buses. Please try again later.");
          console.error("Error fetching buses:", error);
        } finally {
          setLoading(false);
        }
      };
      fetchBuses();
    }
  }, [selectedRoute]);

  const handleRouteChange = (e) => {
    setSelectedRoute(e.target.value); // Actualiza la ruta seleccionada
    setSelectedBus(""); // Reinicia la selección de bus
  };

  const handleBusChange = (e) => {
    setSelectedBus(e.target.value); // Actualiza el bus seleccionado
  };

  const handleSubmit = () => {
    if (selectedRoute && selectedBus) {
      // Navegar a la página de "Payment" con la ruta y el bus seleccionados
      navigate("/payments", { state: { selectedRoute, selectedBus } });
    } else {
      setError("Please select a route and a bus!");
    }
  };

  return (
    <Container className="mt-5">
      <h2 className="text-center mb-4">Select Route and Bus</h2>

      {error && <Alert variant="danger" className="text-center">{error}</Alert>}

      <Row className="justify-content-center">
        <Col md={8}>
          <Card className="shadow">
            <Card.Body>
              <h3 className="text-center mb-4">Select a Route</h3>
              <Form>
                <Form.Group controlId="formRoute" className="mb-4">
                  <Form.Label>Available Routes</Form.Label>
                  <Form.Control
                    as="select"
                    onChange={handleRouteChange}
                    value={selectedRoute}
                    disabled={loading}
                  >
                    <option value="">Choose a Route</option>
                    {Array.isArray(routes) && routes.map((route) => (
                      <option key={route._id} value={route._id}>
                        {route.origin} to {route.destination} - Price: ${route.price}
                      </option>
                    ))}
                  </Form.Control>
                </Form.Group>

                {selectedRoute && (
                  <Form.Group controlId="formBus" className="mb-4">
                    <Form.Label>Available Buses</Form.Label>
                    <Form.Control
                      as="select"
                      onChange={handleBusChange}
                      value={selectedBus}
                      disabled={loading || !buses.length}
                    >
                      <option value="">Choose a Bus</option>
                      {Array.isArray(buses) && buses.map((bus) => (
                        <option key={bus.id} value={bus.id}>
                          {bus.name} - Seats: {bus.capacity}
                        </option>
                      ))}
                    </Form.Control>
                  </Form.Group>
                )}

                <div className="text-center">
                  {loading ? (
                    <Spinner animation="border" role="status">
                      <span className="visually-hidden">Loading...</span>
                    </Spinner>
                  ) : (
                    <Button
                      variant="primary"
                      onClick={handleSubmit}
                      disabled={!selectedRoute || !selectedBus || loading}
                      size="lg"
                    >
                      Continue to Payment
                    </Button>
                  )}
                </div>
              </Form>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Tickets;