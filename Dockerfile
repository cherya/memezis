FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o bin/memezis cmd/memezis/main.go

RUN ls bin

CMD ["/app/bin/memezis"]

EXPOSE 8080