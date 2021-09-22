from google.cloud import pubsub_v1

import requests
import json


class Reader():

    def __init__(self):
        self.array = []
        self.contador = 0
        self.URL = 'http://localhost:5000'

    def getOne(self):
        lengh = len(self.array)
        """self.contador += 1"""

        if lengh > 0:
            return self.array.pop(lengh - 1)
        else:
            print("-> Reader: No hay mas datos en el archivo")

    def load(self):
        print("-> Reader: Iniciando la carga de datos desde el arvhivo")

        try:
            with open("prueba.json", 'r') as data_file:
                self.array = json.loads(data_file.read())

            print(f"-> Reader: Se han cargado correctamente {len(self.array)} datos")
        except Exception as e:
            print(f'-> Reader: No se cargaron los datos, {e}')

    def IniciarCarga(self):
        x = requests.get(self.URL + '/iniciarCarga/python')
        sms = x.json()
        print(sms['message'])

    def sendData(self, mydata):
        x = requests.post(self.URL + '/publicar/python', json = mydata)
        sms = x.json()
        print(sms['message'])

    def FinCarga(self, data):

        x = requests.post(self.URL + '/finalizarCarga/python', json = data)
        sms = x.json()
        print(sms['message'])


"""class LoadTester():"""
