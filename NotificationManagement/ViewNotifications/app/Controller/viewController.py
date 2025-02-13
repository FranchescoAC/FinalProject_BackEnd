from flask import jsonify
from ..Model.viewModel import get_notifications  # Corrección en la importación
import logging

logging.basicConfig(level=logging.DEBUG)

def view_notifications():
    try:
        logging.info("Entrando a la función view_notifications()")
        notifications = get_notifications()
        logging.info(f"get_notifications() devolvió {len(notifications)} notificaciones")
        return jsonify(notifications), 200
    except Exception as e:
        logging.error(f"Error en la función view_notifications(): {e}")
        return jsonify({"error": "Error al obtener notificaciones"}), 500
