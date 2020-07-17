FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN pwd
RUN ls -la

RUN make build

CMD ["/app/bin/memezis", "--env=production.env"]