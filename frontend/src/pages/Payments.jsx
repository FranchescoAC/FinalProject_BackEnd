import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { Container, Card, Alert, Spinner, ListGroup, Row, Col, Button } from "react-bootstrap";
import axios from "axios"; // ✅ Importar Axios para hacer la petición a la API

const Payments = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { selectedRoute, selectedBus } = location.state || {};

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [routeDetails, setRouteDetails] = useState(null);
  const [busDetails, setBusDetails] = useState(null);

  // ✅ Verifica si los datos llegaron correctamente
  useEffect(() => {
    console.log("Datos recibidos en Payments:", selectedRoute, selectedBus);

    if (!selectedRoute || !selectedBus) {
      console.warn("No se recibieron datos, redirigiendo a Tickets...");
      navigate("/tickets");
    } else {
      setRouteDetails(selectedRoute);
      setBusDetails(selectedBus);
      setLoading(false);
    }
  }, [selectedRoute, selectedBus, navigate]);

  // ✅ Función para actualizar la disponibilidad del bus en la base de datos
  const updateBusAvailability = async () => {
    try {
      if (!busDetails) return;
      const newAvailability = busDetails.availability - 1; // Reducir en 1
      if (newAvailability < 0) {
        console.error("Error: No hay disponibilidad en el bus.");
        return;
      }

      console.log("Actualizando disponibilidad del bus...");

      await axios.put(`http://localhost:5000/api/buses/${busDetails.id}`, {
        availability: newAvailability,
      });

      console.log("Disponibilidad del bus actualizada correctamente.");
    } catch (error) {
      console.error("Error al actualizar la disponibilidad del bus:", error);
    }
  };

  // ✅ Cargar PayPal solo cuando los datos están listos
  useEffect(() => {
    if (!routeDetails || !busDetails) return;

    if (document.getElementById("paypal-button-container")) {
      const script = document.createElement("script");
      script.src = `https://www.paypal.com/sdk/js?client-id=ASgyMTqWqImaw3E88hdEwthp8KIkSICwcWhZuoGXXqT5FrVrkrqicxeEm0fsXcge7XupCjLG1vw4HLcC&currency=USD`;
      script.async = true;
      script.onload = () => {
        window.paypal
          .Buttons({
            createOrder: (data, actions) => {
              return actions.order.create({
                purchase_units: [
                  {
                    description: `Bus Ticket for Route ${routeDetails.origin} to ${routeDetails.destination}`,
                    amount: {
                      value: routeDetails.price.toString(),
                      currency_code: "USD",
                    },
                  },
                ],
              });
            },
            onApprove: async (data, actions) => {
              return actions.order.capture().then(async (details) => {
                alert("Payment completed by " + details.payer.name.given_name);
                
                // ✅ Actualizar disponibilidad del bus
                await updateBusAvailability();
                
                navigate("/success");
              });
            },
            onError: (err) => {
              console.error("PayPal error:", err);
              setError("Payment failed. Please try again.");
            },
          })
          .render("#paypal-button-container");
      };
      document.body.appendChild(script);

      return () => {
        document.body.removeChild(script);
      };
    }
  }, [routeDetails, busDetails, navigate]);

  return (
    <Container className="mt-5">
      <h2 className="text-center mb-4">Complete Your Payment</h2>

      {error && <Alert variant="danger" className="text-center">{error}</Alert>}

      <Row className="justify-content-center">
        <Col md={8}>
          <Card className="shadow">
            <Card.Body>
              <h3 className="text-center mb-4">Payment Details</h3>

              {loading ? (
                <div className="text-center">
                  <Spinner animation="border" role="status">
                    <span className="visually-hidden">Loading...</span>
                  </Spinner>
                </div>
              ) : (
                <ListGroup variant="flush">
                  <ListGroup.Item>
                    <strong>Route:</strong> {routeDetails?.origin} to {routeDetails?.destination} (ID: {routeDetails?._id})
                  </ListGroup.Item>
                  <ListGroup.Item>
                    <strong>Bus:</strong> {busDetails?.name} (ID: {busDetails?.id})
                  </ListGroup.Item>
                  <ListGroup.Item>
                    <strong>Price:</strong> ${routeDetails?.price}
                  </ListGroup.Item>
                  <ListGroup.Item>
                    <strong>Ticket:</strong> {busDetails?.availability} {/* ✅ Muestra la disponibilidad actual */}
                  </ListGroup.Item>
                </ListGroup>
              )}

              <div className="mt-4" id="paypal-button-container"></div>

              <div className="mt-3 text-center">
                <Button variant="secondary" onClick={() => navigate("/dashboard")}>Cancel</Button>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Payments;
