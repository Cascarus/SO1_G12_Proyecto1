import json
from  random  import random,randrange
from sys  import getsizeof
from locust import HttpUser, task, between


debug = True



def printDebug(msg):
    if debug:
         print(msg)



class Reader():

    def __init__(self):
        self.array = []
        
    def pickRandom(self):
        length = len(self.array)
        
        if (length > 0):
            random_index = randrange(0, length - 1) if length > 1 else 0
            return self.array.pop(random_index)
        else:
            print (">> Reader: No hay más valores para leer en el archivo.")
            return None
    
    # Cargar el archivo de datos json
    def load(self):
        # Mostramos en consola que estamos a punto de cargar los datos
        print (">> Reader: Iniciando con la carga de datos")
        # Ya que leeremos un archivo, es mejor realizar este proceso con un Try Except
        try:
            # Asignamos el valor del archivo traffic.json a la variable data_file
            with open("prueba.JSON", 'r') as data_file:
                # Con el valor que leemos de data_file vamos a cargar el array con los datos
                self.array = json.loads(data_file.read())
            # Mostramos en consola que hemos finalizado
            
            print (f'>> Reader: Datos cargados correctamente, {len(self.array)} datos -> {getsizeof(self.array)} bytes.')
        except Exception as e:
            # Imprimimos que no pudimos procesar la carga de datos
            print (f'>> Reader: No se cargaron los datos {e}')



class MessageTraffic(HttpUser):
    # Tiempo de espera entre peticiones
    # En este caso, esperara un tiempo de 0.1 segundos a 0.9 segundos (inclusivo) 
    #  entre cada llamada HTTP
    wait_time = between(0.1, 0.9)

    # Este metodo se ejecutara cada vez que empecemos una prueba
    # Este metodo se ejecutara POR USUARIO (o sea, si definimos 3 usuarios, se ejecutara 3 veces y tendremos 3 archivos)
    def on_start(self):
        print (">> MessageTraffic: Iniciando el envio de tráfico")
        # Iniciaremos nuestra clase reader
        self.reader = Reader()
        # Cargaremos nuestro archivo de datos traffic.json
        self.reader.load()

    # Este es una de las tareas que se ejecutara cada vez que pase el tiempo wait_time
    # Realiza un POST a la dirección del host que especificamos en la página de locust
    # En este caso ejecutaremos una petición POST a nuestro host, enviándole uno de los mensajes que leimos
    @task
    def PostMessage(self):
        # Obtener uno de los valores que enviaremos
        random_data = self.reader.pickRandom()
        
        # Si nuestro lector de datos nos devuelve None, es momento de parar
        if (random_data is not None):
            # utilizamos la funcion json.dumps para convertir un objeto JSON de python
            # a uno que podemos enviar por la web (básicamente lo convertimos a String)
            data_to_send = json.dumps(random_data)
            # Imprimimos los datos que enviaremos
            printDebug (data_to_send)

            # Enviar los datos que acabamos de obtener
            self.client.post("/", json=random_data)

        # En este segmento paramos la ejecución del proceso de creación de tráfico
        else:
            print(">> MessageTraffic: Envio de tráfico finalizado, no hay más datos que enviar.")
            # Parar ejecucion del usuario
            self.stop(True) # Se envía True como parámetro para el valor "force", este fuerza a locust a parar el proceso inmediatamente.

    # Este es una de las tareas que se ejecutara cada vez que pase el tiempo wait_time
    # Realiza un GET a la dirección del host que especificamos en la página de locust
    @task
    def GetMessages(self):      
        # Realizar una peticion para recibir los datos que hemos guardado
        self.client.get("/") 









