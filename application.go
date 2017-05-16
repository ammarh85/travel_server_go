package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "./go"
	"log"
	"net/http"
	"os"
	//"io/ioutil"
	//"github.com/gorilla/mux"
	//"fmt"
	//"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Main entry point into the server
// Sets up CORS configurations to ensure it works across domains
func main() {

	log.Printf("Server started")

	router := sw.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}