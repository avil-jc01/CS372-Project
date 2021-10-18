# CS372-Project
CS372 Web Project


# Building Project

to build this project, run:

`docker build . --tag=cs372project`


be sure that you are executing the above when in the repo's directory.

You will need to rebuild the image and re-run the container if you are
making any changes to the backend. For simplicity's sake, the backend
is any .go source code file.

# Running Project

to run the container, do:

`docker run -it -p 8080:8080 -v ${PWD}/templates:/templates cs372project:latest`

where the following: 
> -it : make it an interactive terminal session	(docker stuff - just do this for this container) 
> -p <host_port>:<ctr_port> : map the port on host_port to the container on ctr_port 
> -v ${PWD}/templates:/templates : map this "volume" or directory (${PWD} the current working directory) to /templates inside the container 
> cs372project:latest : the name of the container we built from the previous step

# Contributing

If you are a CS372 project member, be sure to do a `git pull` to get
the latest remote before making any local changes.

Create a branch for your change by doing `git checkout -b
<branch_name>` , where <branch_name> is representative of your
change. For example, `git checkout -b jca-add-customer`.

When you are done editing files, add them to git with `git add
<filename>`. If you have edited multiple files and you want to include
everything, you can use `git add .` as short hand. Next, commit your
changes with `git commit -m 'your commit message here'`. Lastly, do a
`git push` to update the github remote repo and make changes available
to everyone. 

You may see an error from time to time like the below. Simply copy &
paste the suggested `git push` incantation instead.

>fatal: The current branch <your-new-branch> has no upstream branch.
>To push the current branch and set the remote as upstream, use 
>
> git push --set-upstream origin <your-new-branch>
> 
>[CS372-Project]$ git push --set-upstream origin <your-new-branch>
>Enumerating objects: 4, done.
>Counting objects: 100% (4/4), done.
>Delta compression using up to 4 threads
>Compressing objects: 100% (3/3), done.
>Writing objects: 100% (3/3), 2.54 KiB | 2.54 MiB/s, done.
>Total 3 (delta 1), reused 0 (delta 0), pack-reused 0
> >message of successful git push>
> >link to create pull request for your branch>

# Editing the frontend

You can make any edits to the html templates directly while the
container is running. Because of the -v flag in the docker run
incantation, the files are made available to the container in real-time. 

To make any changes to the 'backend' (go code), you will have to
re-run the docker build command in order to compile the new changes.
