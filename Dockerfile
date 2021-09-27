# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /project

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /cs372app

EXPOSE 8080

CMD [ "/cs372app" ]

