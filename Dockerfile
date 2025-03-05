FROM golang:1.23.0-bullseye

WORKDIR /app

COPY ./src .

RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]

