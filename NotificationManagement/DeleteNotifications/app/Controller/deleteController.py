from flask import jsonify
from ..Model.deleteModel import delete_notification_model

# Eliminar notificación
def delete_notification(notification_id):
    try:
        # Llama a la función delete_notification_model del modelo
        result = delete_notification_model(notification_id)
        
        # Verifica que el resultado sea válido y tenga el atributo 'deleted_count'
        if result and hasattr(result, 'deleted_count') and result.deleted_count > 0:
            return jsonify({"message": f"Notification {notification_id} deleted"}), 200
        else:
            return jsonify({"error": f"Notification {notification_id} not found"}), 404
    except Exception as e:
        return jsonify({"error": f"An error occurred: {str(e)}"}), 500
