from bson.objectid import ObjectId
from ..config.db import get_collection

notifications_collection = get_collection('notifications')

# Insertar notificación
def delete_notification_model(notification_id):
    try:
        object_id = ObjectId(notification_id)  # Convertir a ObjectId
        result = notifications_collection.delete_one({"_id": object_id})
        return result  # Devuelve el resultado de la operación
    except Exception as e:
        print(f"Error: {e}")
        return None  # Retorna None en caso de error