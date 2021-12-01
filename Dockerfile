
#declaring our build container
FROM golang as go-build

#setting our working dir inside container
WORKDIR /cs372

COPY go.* ./

RUN go mod download

#copy repo files into container
COPY . .

#build our binary 
RUN go build -v -o /cs372-project

#using small ubuntu container for server
FROM ubuntu as server

#move the executable from the build container to the server container
COPY --from=go-build /cs372-project /

#copy templates over to ubuntu server container
COPY templates/ /templates/

#copy static resources over
COPY static/ /static/

#run the binary on docker run command
CMD ["/cs372-project"]
