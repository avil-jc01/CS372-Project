# CS372-Project
CS372 Web Project


# Building Project

to build this project, run:

docker build . --tag=cs372project


be sure that you are executing the above when in the repo's directory.

# Running Project

to run the container, do:

docker run -it -p 8080:8080 cs372project:latest

where the following:
-it : make it an interactive terminal session (docker stuff - just do this for this container)
-p <host_port>:<ctr_port> : map the port on host_port to the container on ctr_port
cs372project:latest : the name of the container we built from the previous step


