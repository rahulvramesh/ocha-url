# syntax=docker/dockerfile:1
FROM golang:1.16 as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go mod vendor && CGO_ENABLED=0 GOOS=linux  go build -a -tags netgo -ldflags '-w' -o bin/ocha cmd/ocha/main.go

FROM alpine:latest

COPY --from=builder build/bin/ocha .

RUN chmod +x ocha

CMD ["./ocha"]
