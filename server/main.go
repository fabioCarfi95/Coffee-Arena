package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/fabioCarfi95/Coffee-Arena/server/methods"

	"github.com/gorilla/mux"
)

var (
	help    = flag.Bool("help", false, "Show help")
	address string
	port    int
	DB_USER string
	DB_PWD  string
	DB_HOST string
	DB_PORT int
	DB_NAME string
	db      *sql.DB
)

func handleRequests() {
	// Create new router
	router := mux.NewRouter()
	log.Println("Creating routes")
	// Specify endpoints
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/userinfo", userInfo).Methods("POST")

	log.Printf("Listen&Serve on %s:%d", address, port)
	log.Fatal(http.ListenAndServe(address+":"+strconv.Itoa(port), router))
}

func init() {
	// Bind flags
	flag.StringVar(&address, "address", "localhost", "REST service base address")
	flag.IntVar(&port, "port", 8080, "REST service port")

	flag.StringVar(&DB_USER, "DB_USER", "coffee_arena_admin", "DB service username")
	flag.StringVar(&DB_PWD, "DB_PWD", "secret", "DB service password")
	flag.StringVar(&DB_HOST, "DB_HOST", "localhost", "DB service host")
	flag.IntVar(&DB_PORT, "DB_PORT", 3306, "DB service port")
	flag.StringVar(&DB_NAME, "DB_NAME", "coffee_arena_db", "DB service name")

	// Parse flags
	flag.Parse()

	// Usage demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

}

func main() {
	setupDB()
	handleRequests()
}
