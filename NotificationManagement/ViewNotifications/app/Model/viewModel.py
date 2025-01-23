from bson.objectid import ObjectId
from ..config.db import get_collection

notifications_collection = get_collection('notifications')

# Insertar notificación
def get_notifications():
    notifications = notifications_collection.find({})
    # Convertir cada documento a un formato JSON serializable
    return [{**doc, "_id": str(doc["_id"])} for doc in notifications]
