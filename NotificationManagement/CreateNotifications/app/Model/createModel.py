from bson.objectid import ObjectId
from ..config.db import get_collection

# Collection reference
notifications_collection = get_collection('notifications')

# Insert notification into the database
def create_notification(data):
    result = notifications_collection.insert_one(data)
    return str(result.inserted_id)  # Convert ObjectId to string

