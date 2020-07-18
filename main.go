package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultHost = "0.0.0.0"
const defaultPort = "8080"

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	hostPort := host+":"+port

	fmt.Println("Starting web server on http://"+hostPort)

	http.Handle("/", http.FileServer(http.Dir("./static")))


	log.Fatal(http.ListenAndServe(host+":"+port, nil))

}
