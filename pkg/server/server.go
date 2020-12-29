package server

import (
	"fmt"
	"net/http"
	"os"
)

//Handler - Function for handling server
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service Not Found")
}

//Setport - Set Port for server
func Setport() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
