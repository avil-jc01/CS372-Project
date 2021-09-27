# syntax=docker/dockerfile:1

FROM golang

WORKDIR  /go/src/github.com/avil-jc01/CS372-Project/
COPY .  /go/src/github.com/avil-jc01/CS372-Project/

RUN go mod download
RUN go get -v -t  .

RUN go build -o /cs372app

EXPOSE 8080

CMD [ "/cs372app" ]

