# CS372-Project
CS372 Web Project


# Building Project

to build this project, run:

docker build . --tag=cs372project


be sure that you are executing the above when in the repo's directory.

# Running Project

to run the container, do:

docker run -it -p 8080:8080 -v ${PWD}/templates:/templates cs372project:latest

where the following:

-it : make it an interactive terminal session (docker stuff - just do this for this container)

-p <host_port>:<ctr_port> : map the port on host_port to the container on ctr_port

-v ${PWD}/templates:/templates : map this "volume" or directory (${PWD} the current working directory) to /templates inside the container

cs372project:latest : the name of the container we built from the previous step

# Editing the frontend

You can make any edits to the html templates directly while the
container is running. Because of the -v flag in the docker run
incantation, the files are made available to the container in real-time. 

To make any changes to the 'backend' (go code), you will have to re-run the docker build command in order to compile the new changes
