package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	port := flag.String("port", "8080", "HTTP port")

	flag.Parse()

	//router configuration
	router := mux.NewRouter()

	//serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	fmt.Printf("Starting...")

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
