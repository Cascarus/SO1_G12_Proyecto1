FROM golang:latest

WORKDIR /usr/src/goApp

COPY . ./

RUN apt-get update

RUN go mod download

ENV DB=Olympics
ENV COLLECTION=Tuits
ENV MONGO_CS=mongodb://go-mongo:NWm0Ub0V1DZxVOLZb6IyMPa4a6HCHqWEuj8DZhjHV9VVFScnSWFDk0ky2xX61sZemUeq7Q61Tv4stiJKYVrXNw==@go-mongo.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@go-mongo@
ENV USER=root
ENV PASS=123456
ENV PROYECT=deft-idiom-324423
ENV TOPIC=projects/deft-idiom-324423/topics/olympics
ENV DB_ADDR=34.122.151.115
ENV GIN_MODE=release

#RUN apt-get install -y nodejs

#RUN apt-get install -y npm

#RUN npm install -g nodemon

EXPOSE 8080

#CMD ["nodemon", "--exec", "go", "run", "main.go"]

CMD ["go", "run", "main.go"]

#docker run -idt -p 1200:8010 --name go-api go:api
