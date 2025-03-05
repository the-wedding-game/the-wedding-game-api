FROM golang:1.23.0-alpine AS builder

WORKDIR /app
ADD ./src .

RUN CGO_ENABLED=0 go build -o server

FROM scratch

COPY --from=builder /app/server /app/server

CMD ["/app/server"]
