import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { Container, Card, Alert, Spinner, ListGroup, Row, Col, Button } from "react-bootstrap";


const Payments = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { selectedRoute, selectedBus } = location.state || {};

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [routeDetails, setRouteDetails] = useState(null);
  const [busDetails, setBusDetails] = useState(null);


 

  useEffect(() => {
    if (!selectedRoute || !selectedBus) {
      navigate("/tickets");
    } else {
      setLoading(true);
      setTimeout(() => {
        setRouteDetails({
          id: selectedRoute,
          origin: "Quito",
          destination: "Guayaquil",
          price: 15,
        });
        setBusDetails({
          id: selectedBus,
          name: "Bus trasesmeraldas",
          capacity: 40,
        });
        setLoading(false);
      }, 1000);
    }
  }, [selectedRoute, selectedBus, navigate]);

  useEffect(() => {
    if (routeDetails && busDetails && document.getElementById("paypal-button-container")) {
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
            onApprove: (data, actions) => {
              return actions.order.capture().then((details) => {
                alert("Payment completed by " + details.payer.name.given_name);
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
                    <strong>Route:</strong> {routeDetails?.origin} to {routeDetails?.destination} (ID: {routeDetails?.id})
                  </ListGroup.Item>
                  <ListGroup.Item>
                    <strong>Bus:</strong> {busDetails?.name} (ID: {busDetails?.id})
                  </ListGroup.Item>
                  <ListGroup.Item>
                    <strong>Price:</strong> ${routeDetails?.price}
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
