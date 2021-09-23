FROM golang:latest

WORKDIR /usr/src/proxy

COPY . ./

RUN apt-get update

RUN go get github.com/GoogleCloudPlatform/cloudsql-proxy/cmd/cloud_sql_proxy

EXPOSE 1443

CMD ["echo", "Hola Mundo!"]
#CMD ["cloud_sql_proxy", "-instances==tcp:0.0.0.0:1433", "-credential_file=/usr/src/proxy/sql/k.json"]