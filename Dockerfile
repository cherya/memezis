FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

RUN ls

CMD ["/app/bin/memezis"]

EXPOSE 8080