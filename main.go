package main

import (
	"fmt"
	"log"
	"net/http"

	Server "github.com/Andromeda/NotFound/pkg/server"
)

func main() {
	http.HandleFunc("/", Server.Handler)
	port := Server.Setport()

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
