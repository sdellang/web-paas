package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdellang/charta"
)

func main() {

	port := flag.String("port", "8080", "HTTP port")

	flag.Parse()

	//router configuration
	router := mux.NewRouter()

	router.HandleFunc("/api/cluster", charta.GetClusterView).Methods("GET")
	router.HandleFunc("/api/namespaces", charta.GetNamespaces).Methods("GET")
	router.HandleFunc("/api/namespaces/{name}", charta.GetNamespaceView).Methods("GET")
	//serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/dist/")))
	fmt.Printf("Starting...")

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
