# syntax=docker/dockerfile:1

FROM golang as go-build
ENV GO111MODULE=on
WORKDIR /usr/local/go/src/CS372-Project

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY models/ models/
COPY handlers/ handlers/
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /cs372-project

FROM ubuntu as server

COPY --from=go-build /cs372-project /


CMD ["/cs372-project"]