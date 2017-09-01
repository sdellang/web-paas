package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	/*fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}*/

	if r.URL.Path == "/pizza.jpg" {
		bytes, err := ioutil.ReadFile("/web/pizza.jpg") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
		if _, err := w.Write(bytes); err != nil {
			log.Println("unable to write image.")
		}
	} else {
		bytes, err := ioutil.ReadFile("/web/index.html") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		//fmt.Println(bytes) // print the content as 'bytes'
		output := string(bytes) // convert content to a 'string'
		fmt.Fprintf(w, output)  // print the content as a 'string'

		//fmt.Fprintf(w, "Hello Alex!") // send data to client side
	}
}

func main() {

	log.Println("I'll try listening on the following addresses:")
	ifaces, errNet := net.Interfaces()
	// handle err
	if errNet != nil {
		log.Fatal("net.Interfaces(): ", errNet)
	} else {
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			// handle err
			if err != nil {
				log.Fatal("ifaces.Addrs(): ", err)
			} else {
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}
					log.Println("http://" + ip.String())
				}
			}
		}
	}

	http.HandleFunc("/", sayhelloName)     // set router
	err := http.ListenAndServe(":80", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
