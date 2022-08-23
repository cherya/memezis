FROM golang:1.18-alpine

RUN apk add --no-cache make

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN pwd
RUN ls -la

RUN make build

CMD ["/app/bin/memezis", "--env=production.env"]
