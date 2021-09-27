package main

import (
    	"github.com/avil-jc01/CS372-Project/handlers"
    	"database/sql"
    	"log"
    	_ "github.com/mattn/go-sqlite3"
    	"net/http"
)

func main() {
	
    	log.Println("D: Starting - Running CS372 App on port 8080")
	
    	//create db connection
    	database, err := sql.Open("sqlite3", "./budget-app.db")
    	if err != nil {
    		panic(err)
    	}

	defer database.Close()
	
    	//creating a new http multiplexer
    	mux := http.NewServeMux()

    	//add fileserver for static resources
    	//fileserver := http.FileServer(http.Dir("./static/"))

    	//adding handler functions for each handler
    	hh := http.HandlerFunc(handlers.HomeHandler)
    	
    	//use the mux to handle requests to pages
    	mux.Handle("/", hh)

    	//start the webserver
    	http.ListenAndServe(":8080", mux)

    }

