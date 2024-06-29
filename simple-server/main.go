package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "Port number on which to run simple http server")
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)

	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalln(err)
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	log.Println("Received a request at / route")
	fmt.Fprintf(w, "You have reached a simple web server...")
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("Received a request at /hello route")
	fmt.Fprintf(w, "Hello there!")
}
