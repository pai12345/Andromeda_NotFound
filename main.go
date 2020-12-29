package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//handler - Function for handling server
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service Not Found")
}

//setport - Set Port for App
func setport() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	http.HandleFunc("/", handler)
	port := setport()
	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
