package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	//sw "./go"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {

	log.Printf("Server started")

	//router := sw.NewRouter()
	



	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	//log.Fatal(http.ListenAndServe(":"+port, router))


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":"+port, router))



	f, _ := os.Create("/var/log/golang/golang-server.log")
	defer f.Close()
	log.SetOutput(f)

	const indexPage = "public/index.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if buf, err := ioutil.ReadAll(r.Body); err == nil {
				log.Printf("Received message: %s\n", string(buf))
			}
		} else {
			log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
			http.ServeFile(w, r, indexPage)
		}
	})

	http.HandleFunc("/scheduled", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			log.Printf("Received task %s scheduled at %s\n", r.Header.Get("X-Aws-Sqsd-Taskname"), r.Header.Get("X-Aws-Sqsd-Scheduled-At"))
		}
	})

	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}