from bson.objectid import ObjectId
from ..config.db import get_collection
import logging

logging.basicConfig(level=logging.DEBUG)

def get_notifications():
    try:
        notifications_collection = get_collection('notification_management', 'notifications')
        logging.info("Conexión a la colección 'notifications' exitosa.")

        # Convertir directamente en una lista
        notifications = [{**doc, "_id": str(doc["_id"])} for doc in notifications_collection.find({})]
        
        logging.info(f"Se encontraron {len(notifications)} notificaciones.")
        return notifications
    except Exception as e:
        logging.error(f"Error al obtener notificaciones: {e}")
        return []  # Devuelve una lista vacía en caso de error
