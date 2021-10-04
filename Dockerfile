# syntax=docker/dockerfile:1

#We are using golang as our base image
FROM golang

#Golang has a specific dir heirarchy - we declare that as our workdir
WORKDIR  /usr/local/go/src/github.com/avil-jc01/CS372-Project/

#Copy files from "here" (github repo dir) to the WORKDIR in the container
COPY .  /usr/local/go/src/github.com/avil-jc01/CS372-Project/

#Use Go Mod tool to download & get our dependencies
RUN go mod download
RUN go get -v -t  .

#Build our linux elf - the binary file we run
RUN GO111MODULE=on  CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o /cs372app

#Listen for http requests on port 8080
EXPOSE 8080

#When container is started, run this command (the binary from above)
CMD [ "/cs372app" ]

