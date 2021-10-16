# syntax=docker/dockerfile:1

FROM golang as go-build
WORKDIR /usr/local/go/src/CS372-Project

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY models/ models/
COPY handlers/ handlers/
COPY main.go .

RUN cd /usr/local/go/src/CS372-Project && go build -v -o /cs372-project

FROM ubuntu as server

COPY --from=go-build /cs372-project /


CMD ["/cs372-project"]