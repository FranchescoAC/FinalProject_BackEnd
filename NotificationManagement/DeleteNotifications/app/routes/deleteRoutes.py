from flask import Blueprint
from ..Controller.deleteController import delete_notification


notification_routes = Blueprint('notification_routes', __name__)

notification_routes.route('/notifications/<string:notification_id>', methods=['DELETE'])(delete_notification)