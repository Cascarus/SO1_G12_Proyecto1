from flask import Flask, request, jsonify
from pymongo import MongoClient
from google.cloud import pubsub_v1

import pyodbc

import json
from google.auth import jwt

service_account_info = json.load(open("PS.json"))

app = Flask(__name__)

'''
Conexion con Cosmos
'''
db_name = "Olympics"
host = "go-mongo.mongo.cosmos.azure.com"
port = 10255
username = "go-mongo"
password = "NWm0Ub0V1DZxVOLZb6IyMPa4a6HCHqWEuj8DZhjHV9VVFScnSWFDk0ky2xX61sZemUeq7Q61Tv4stiJKYVrXNw=="
args = "ssl=true&retrywrites=false&ssl_cert_reqs=CERT_NONE"

connection_uri = f"mongodb://{username}:{password}@{host}:{port}/{db_name}?{args}"

client = MongoClient(connection_uri)
cosmosDB = client[db_name]
container = cosmosDB['Tuits']

'''
Conexion con Google SQL SERVER
'''
server = '35.192.164.59'
database = 'Olympics'
username = 'sqlserver'
password = '123456'
conSQL = pyodbc.connect(
    'DRIVER={ODBC Driver 17 for SQL Server};'
    'SERVER=' + server +
    ';DATABASE=' + database +
    ';UID=' + username +
    ';PWD=' + password
)

cursor = conSQL.cursor()
cursor.execute("SET DATEFORMAT dmy")

cargar: bool = False
contadorSQL: int = 0
contadorCosmos: int = 0

project_id = "deft-idiom-324423"
topic_id = "olympics"


@app.route('/iniciarCarga/python', methods=['GET'])
def iniciarCarga():
    global cargar, contadorSQL, contadorCosmos, idCosmos

    if cargar == False:
        cargar = True
        contadorSQL = 0
        contadorCosmos = 0
        return jsonify({'message': 'Se ha iniciado la conexion correctamente, puede enviar datos'})

    return jsonify({'message': 'Ya hay una conexion iniciada'})


@app.route('/publicar/python', methods=['POST'])
def publicar():
    global cargar
    if not cargar:
        return jsonify({'message': 'Debe iniciar una conexion para poder ingresar datos a la DB'})
    body = request.get_json()

    publicarSQL(body)
    publicarCosmos(body)
    return jsonify({'message': 'Se ha cargado el dato exitosamente a la DB'})


@app.route('/finalizarCarga/python', methods=['POST'])
def finCarga():
    global cargar, contadorSQL, contadorCosmos
    cargar = False

    body = request.get_json()
    body["guardados"] = contadorSQL
    body["api"] = "Python"
    body["db"] = "Cloud SQL"
    '''Enviar datoas a Google PUB/SUB'''
    publicarPub(body)
    body["guardados"] = contadorCosmos
    body["db"] = "CosmosDB"
    publicarPub(body)
    return jsonify({'message': 'Se ha finalizado la carga de datos exitosamente!'})


def publicarSQL(body):
    listaHastag = body["hashtags"]
    hashtags = ""

    for hasht in listaHastag:
        hashtags += hasht + ','

    hashtags = hashtags[0:len(hashtags) - 1]

    query = """INSERT INTO OLIMPIC(nombre, comentario, fecha, hashtags, upvotes, downvotes)
                   VALUES('{0}','{1}','{2}','{3}',{4},{5})""".format(body["nombre"], body["comentario"], body["fecha"],
                                                                     hashtags, body["upvotes"], body["downvotes"])
    try:
        cursor.execute(query)
        conSQL.commit()
        incrementSqlCounter()
    except:
        return


def publicarCosmos(body):
    try:
        container.insert_one(body)
        incrementCosmosCounter()
        return jsonify({'respuesta': 'Todo bien, nada mal'})
    except:
        return jsonify({'respuesta': 'Todo mal, nada bien'})


def publicarPub(data):
    publisher = pubsub_v1.PublisherClient()
    topic_path = publisher.topic_path(project_id, topic_id)
    data = json.dumps(data)
    data = data.encode("utf-8")
    future = publisher.publish(topic_path, data)
    sms = future.result()
    print(f"Published {data} to {topic_path}: {sms}")


def getIndex():
    global idCosmos
    idCosmos += 1
    return idCosmos


def decrementIndex():
    global idCosmos
    idCosmos -= 1


def incrementSqlCounter():
    global contadorSQL
    contadorSQL += 1


def incrementCosmosCounter():
    global contadorCosmos
    contadorCosmos += 1


if __name__ == "__main__":
    """app.run(debug=True)"""
    from waitress import serve

    print("server on port", 5000)
    serve(app, port=5000)

