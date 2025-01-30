FROM golang:1.21.5 AS builder

WORKDIR /app

COPY . .

RUN make build

FROM docker.io/alpine:3.17 AS target
ENV PORT=8080
COPY --from=builder /app/build/synmail-api /usr/local/bin/synmail-api
ENTRYPOINT []
CMD ["/usr/local/bin/synmail-api"]