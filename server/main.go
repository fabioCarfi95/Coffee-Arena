package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	help    = flag.Bool("help", false, "Show help")
	address string
	port    int
)

func handleRequests() {
	// Create new router
	router := mux.NewRouter()
	log.Println("Creating routes")
	
	// Specify endpoints
	router.HandleFunc("/", homePage).Methods("GET")

	log.Printf("Listen&Serve on %s:%d", address, port)
	log.Fatal(http.ListenAndServe(address+":"+strconv.Itoa(port), router))
}

func init() {
	// Bind flags
	flag.StringVar(&address, "address", "0.0.0.0", "REST service base address")
	flag.IntVar(&port, "port", 10000, "REST service port")

	// Parse flags
	flag.Parse()

	// Usage demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	handleRequests()
}
