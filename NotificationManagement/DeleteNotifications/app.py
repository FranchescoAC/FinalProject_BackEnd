from flask import Flask
from app.routes.deleteRoutes import notification_routes

app = Flask(__name__)

app.register_blueprint(notification_routes)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)

