from flask import Blueprint
from ..Controller.createController import send_notification


notification_routes = Blueprint('notification_routes', __name__)

notification_routes.route('/notifications', methods=['POST'])(send_notification)