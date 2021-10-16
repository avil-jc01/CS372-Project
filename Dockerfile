# syntax=docker/dockerfile:1

#declaring our build container
FROM golang as go-build

#setting our working dir inside container
WORKDIR /cs372

#fetch our dependencies - these will cache
RUN go get github.com/mattn/go-sqlite3

#copy repo files into container
COPY . .

#build our binary 
RUN  go build -v -o /cs372-project

#using small ubuntu container for server
FROM ubuntu as server

#move the executable from the build container to the server container
COPY --from=go-build /cs372-project /

#run the binary on docker run command
CMD ["/cs372-project"]