from flask import Flask
from app.routes.viewRoutes import notification_routes

app = Flask(__name__)

# Registrar rutas
app.register_blueprint(notification_routes)

if __name__ == '__main__':
    app.run(port=5003, debug=True)
