from flask import Flask
from app.routes.viewRoutes import notification_routes  # Corrección en la importación

app = Flask(__name__)

app.register_blueprint(notification_routes)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5003, debug=True)
