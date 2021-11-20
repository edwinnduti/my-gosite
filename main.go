package main

import (
	"log"
	"net/http"
	"os"

	"github.com/edwinnduti/my-gosite/router"
)

//credentials
var (
	TO_EMAIL = []string{os.Getenv("USERNAME")}
	PASSWORD = os.Getenv("PASSWORD")
	USERNAME = os.Getenv("USERNAME")
	Port     = os.Getenv("PORT")
)

func main() {
	// call router
	r := router.Route()

	// set port
	if Port == "" {
		Port = "8081"
	}

	//start server
	server := &http.Server{
		Handler: r,
		Addr:    ":" + Port,
	}

	//log output
	log.Printf("Listening on port: %v", Port)
	server.ListenAndServe()
}
