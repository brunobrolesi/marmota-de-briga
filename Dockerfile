FROM golang:1.22-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN go build -C ./cmd -buildvcs=false -o bin/marmota-de-briga

FROM alpine:3.19

EXPOSE 8080

COPY --from=builder /usr/src/app/cmd/bin .

ENTRYPOINT ["./marmota-de-briga"] 