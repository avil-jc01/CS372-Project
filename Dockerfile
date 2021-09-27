# syntax=docker/dockerfile:1

FROM golang

WORKDIR  /usr/local/go/src/github.com/avil-jc01/CS372-Project/
COPY .  /usr/local/go/src/github.com/avil-jc01/CS372-Project/

RUN go mod download
RUN go get -v -t  .

RUN GO111MODULE=on  CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o /cs372app

EXPOSE 8080

CMD [ "/cs372app" ]

