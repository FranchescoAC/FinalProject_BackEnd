from pymongo import MongoClient

# Conectar a MongoDB
client = MongoClient("mongodb://3.92.175.95:27017/")  # Cambia el URI según tus necesidades
db = client['notification_management']  # Nombre de la base de datos

def get_collection(collection_name):
    return db[collection_name]
