FROM golang:1.7
RUN apt-get update && apt-get install -y uuid-runtime && apt-get install -y net-tools
RUN go get github.com/gocql/gocql
EXPOSE 8443
RUN mkdir /go/src/iAccounts
WORKDIR /go/src/iAccounts
ADD . /go/src/iAccounts
RUN go build main/webServer.go
CMD ["./webServer", "8443"]
