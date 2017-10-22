package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type StatusRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func main() {

	port := flag.String("port", "80", "HTTP port")

	flag.Parse()

	//router configuration
	router := mux.NewRouter()

	//serving static files
	router.PathPrefix("/").Handler(wrapHandler(http.FileServer(http.Dir("./web/"))))
	fmt.Printf("Starting... \n")

	log.Println("I'll try listening on the following addresses:")
	ifaces, errNet := net.Interfaces()
	// handle err
	if errNet != nil {
		log.Fatal("net.Interfaces(): ", errNet)
	} else {
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			// handle err
			if err == nil {
				log.Fatal("ifaces.Addrs(): ", err)
			} else {
				for _, addr := range addrs {
					v := addr.(type)
					log.Println("http://" + v.IP.String())
				}
			}
		}
	}

	log.Fatal(http.ListenAndServe(":"+*port, router))
}

func (w *StatusRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	w.ResponseWriter.WriteHeader(status)
}

func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srw := &StatusRespWr{ResponseWriter: w}
		h.ServeHTTP(srw, r)
		if srw.status >= 400 { // 400+ codes are the error codes
			log.Printf("Error status code: %d when serving path: %s",
				srw.status, r.RequestURI)
		}
	}
}
