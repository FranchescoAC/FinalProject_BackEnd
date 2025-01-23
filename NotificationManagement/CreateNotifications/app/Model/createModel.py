from bson.objectid import ObjectId
from ..config.db import get_collection

notifications_collection = get_collection('notifications')

# Insertar notificación
def create_notification(data):
    result = notifications_collection.insert_one(data)
    return str(result.inserted_id)  # Convertir ObjectId a cadena
