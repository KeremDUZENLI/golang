FROM golang

RUN mkdir /+SQLC

ADD . /+SQLC

WORKDIR /+SQLC

RUN go build -o main.go

EXPOSE 8080

CMD ["/+SQLC/main"]
