from flask import Blueprint
from ..Controller.viewController import view_notifications  # Corrección en la importación

notification_routes = Blueprint('notification_routes', __name__)

notification_routes.route('/notifications', methods=['GET'])(view_notifications)
