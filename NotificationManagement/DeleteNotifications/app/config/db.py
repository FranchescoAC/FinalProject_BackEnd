from pymongo import MongoClient

# Conectar a MongoDB
client = MongoClient("mongodb://ec2-54-84-63-64.compute-1.amazonaws.com:27017/")  # Cambia el URI según tus necesidades
db = client['notification_management']  # Nombre de la base de datos

def get_collection(collection_name):
    return db[collection_name]
