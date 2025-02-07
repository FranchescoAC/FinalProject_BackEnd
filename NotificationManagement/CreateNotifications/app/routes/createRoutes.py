from flask import Blueprint
from ..Controller.createController import send_notification
from ..Controller.webhookController import receive_webhook

notification_routes = Blueprint('notification_routes', __name__)

# Route to create notifications manually (if needed)
notification_routes.route('/notifications', methods=['POST'])(send_notification)

# Route to receive webhooks from UserManagement
notification_routes.route('/api/notifications/new-user', methods=['POST'])(receive_webhook)
