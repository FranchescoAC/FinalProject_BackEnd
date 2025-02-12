import os
from pymongo import MongoClient
from pymongo.errors import ConnectionFailure, OperationFailure

MONGODB_URI = os.getenv("MONGODB_URI")

if not MONGODB_URI:
    raise ValueError("La variable de entorno MONGODB_URI no está definida.")

print(f"MONGODB_URI: {MONGODB_URI}")

try:
    client = MongoClient(MONGODB_URI)  # Mantener la conexión abierta
    client.admin.command('ping')  # Verificar conexión
    print("Conexión a MongoDB exitosa.")
except (ConnectionFailure, OperationFailure) as e:
    print(f"Error al conectar a MongoDB: {e}")
    raise

def get_collection(database_name, collection_name):
    db = client[database_name]
    return db[collection_name]
