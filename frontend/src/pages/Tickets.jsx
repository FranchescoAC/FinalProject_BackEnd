import { useEffect, useState } from "react";
import { getRoutes, getBuses } from "../services/ticketService";
import { Container, Row, Col, Button, Card, Form, Alert, Spinner } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

const Tickets = () => {
  const [routes, setRoutes] = useState([]);
  const [buses, setBuses] = useState([]);
  const [selectedRoute, setSelectedRoute] = useState("");
  const [selectedBus, setSelectedBus] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  // Mapeo de origen a nombre de bus permitido
  const busFilterMapping = {
    "Quito": "Bus Fénix",
    "Guayaquil": "Bus del Sol",
    "Cuenca": "Bus Andino",
    "Loja": "Bus Loja Express",
    "Machala": "Bus del Pacífico",
    "Manta": "Bus Marina",
    "Portoviejo": "Bus Costa",
    "Santo Domingo": "Bus Tropical",
    "Ibarra": "Bus Sierra",
    "Tulcán": "Bus Frontera",
    "Riobamba": "Bus Volcán",
    "Ambato": "Bus Central",
    "Latacunga": "Bus Latacunga",
    "Tena": "Bus Amazónico",
    "Puyo": "Bus Selva",
    "Macas": "Bus Oriente"
  };

  // Obtener rutas al cargar la vista
  useEffect(() => {
    const fetchRoutes = async () => {
      setLoading(true);
      try {
        const routesData = await getRoutes();
        setRoutes(routesData);
      } catch (error) {
        setError("Error fetching routes. Please try again later.");
        console.error("Error fetching routes:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchRoutes();
  }, []);

  // Obtener y filtrar los buses cuando se selecciona una ruta
  useEffect(() => {
    if (selectedRoute) {
      const fetchBuses = async () => {
        setLoading(true);
        try {
          const busesData = await getBuses();
          // Encontrar el objeto de la ruta seleccionada para obtener su origen
          const selectedRouteObj = routes.find(route => route._id === selectedRoute);
          if (selectedRouteObj) {
            const allowedBusName = busFilterMapping[selectedRouteObj.origin];
            // Si existe un bus permitido para ese origen, filtrar los buses
            const filteredBuses = allowedBusName 
              ? busesData.filter(bus => bus.name === allowedBusName)
              : [];
            setBuses(filteredBuses);
          } else {
            setBuses([]);
          }
        } catch (error) {
          setError("Error fetching buses. Please try again later.");
          console.error("Error fetching buses:", error);
        } finally {
          setLoading(false);
        }
      };
      fetchBuses();
    }
  }, [selectedRoute, routes]);

  const handleRouteChange = (e) => {
    setSelectedRoute(e.target.value);
    setSelectedBus("");
  };

  const handleBusChange = (e) => {
    setSelectedBus(e.target.value);
  };

  const handleSubmit = () => {
    if (selectedRoute && selectedBus) {
      // Buscar la ruta y el bus completos antes de navegar
      const selectedRouteObj = routes.find(route => route._id === selectedRoute);
      const selectedBusObj = buses.find(bus => String(bus.id) === String(selectedBus)); // 🔥 Convertir ambos a string
  
      console.log("Navegando a Payments con datos:", selectedRouteObj, selectedBusObj); // 🔍 Verificar en consola
  
      if (selectedRouteObj && selectedBusObj) {
        navigate("/payments", { state: { selectedRoute: selectedRouteObj, selectedBus: selectedBusObj } });
      } else {
        setError("Error: No se pudo encontrar la información seleccionada.");
      }
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
