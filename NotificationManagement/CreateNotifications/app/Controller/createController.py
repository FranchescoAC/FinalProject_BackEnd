from flask import jsonify, request
from ..Model.createModel import create_notification

# Crear notificación
def send_notification():
    data = request.get_json()
    notification_id = create_notification(data)
    return jsonify({"message": "Notification created", "id": str(notification_id)}), 201
