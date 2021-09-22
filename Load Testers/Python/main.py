import time
import random
from TestLoader import Reader

if __name__ == "__main__":

    app = Reader()
    app.load()
    app.IniciarCarga()
    contador = 0
    conta_tiempo = 0

    while(len(app.array) > 0):
        contador += 1
        dato = app.getOne()
        app.sendData(dato)
        tiempo = round(random.uniform(0,0.4), 1)
        conta_tiempo += tiempo
        time.sleep(tiempo)

    print("-> Load Tester: se han enviado todos los datos")

    data = {"guardados": 0, "api": "Python", "tiempoDeCarga": conta_tiempo, "db": ""}

    app.FinCarga(data)
    """https://www.w3schools.com/python/ref_requests_post.asp"""