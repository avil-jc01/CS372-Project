package main
// golang programs begin with a package declaration - C++ -like


// import statement for dependencies
import (
	// we are using github.com as the "root" of our project
    	"CS372-Project/handlers"
    	"database/sql"
    	"log"
    	_ "github.com/mattn/go-sqlite3"
    	"net/http"
)

// golang programs have main() as entrypoint - C -like
func main() {

	// log prints to stdout - the docker container captures this output
    	log.Println("D: Starting - Running CS372 App on port 8080")
	
    	//create db connection
    	database, err := sql.Open("sqlite3", "./cs372-app.db")
    	if err != nil {
    		panic(err)
    	}

	//close databse connection before main() exits
	defer database.Close()
	
    	//http requests are routed to the endpoint with a multiplexer
    	mux := http.NewServeMux()

    	//adding handler functions - the handler defines the GET/POST/etc. action
	// see ./handlers/homehandler.go
    	hh := http.HandlerFunc(handlers.HomeHandler)
    	aah := http.HandlerFunc(handlers.AddAutoHandler)

	//attach the handler to the multiplexer with the appropriate endpoint
    	mux.Handle("/", hh)
	mux.Handle("/add-auto", aah)

	//some handlers we will probably be building in the future:
	//mux.Handle("/view-autos", nil)

	//mux.Handle("/update-auto", nil)
	//mux.Handle("/login", nil)
	
    	//start the webserver
    	http.ListenAndServe(":8080", mux)

    }

