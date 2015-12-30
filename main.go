package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
)

// redirect function does the redirection
func redirect(w http.ResponseWriter, r *http.Request, location string) {
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusFound)
}

func main() {
	// Handle cmdline arguments
	location := flag.String("location", "localhost:80", "The URI at which the ")
	port := flag.Int("port", 80, "Port at which the http server binds to")
	flag.Parse()

	// Validate
	isURL := govalidator.IsURL(*location)
	if !isURL {
		log.Fatal("The location you provided is not a valid URL")
	}
	if *port < 1 || *port > 65536 {
		log.Fatal("Port number specified is invalid")
	}

	log.Printf("Starting web server... Location: %s, Port: %d", *location, *port)

	// Start redirection service at path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect(w, r, *location)
	})
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatal("Error starting web server. ", err)
	}
}
