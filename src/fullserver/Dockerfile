FROM golang:1.14

WORKDIR /src

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go get -u github.com/gorilla/mux

EXPOSE 8080

RUN go build server.go

CMD ["server"]