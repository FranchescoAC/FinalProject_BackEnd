from flask import jsonify, request
from ..Model.viewModel import get_notifications

# Crear notificación
def view_notifications():
    notifications = get_notifications()
    return jsonify(notifications), 200