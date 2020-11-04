package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cardil/homedir/internal/view/controller"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	http.HandleFunc("/", controller.List)

	log.Printf("Listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
