FROM golang:latest

WORKDIR /usr/src/goApp

COPY . ./

RUN apt-get update

#RUN mod init goApi && go get

RUN go get -u github.com/gin-gonic/gin

#RUN apt-get install -y nodejs

#RUN apt-get install -y npm

#RUN npm install -g nodemon

EXPOSE 8080

#CMD ["nodemon", "--exec", "go", "run", "main.go"]

CMD ["go", "run", "main.go"]


#docker run -idt -p 1200:8010 --name go-api go:api
