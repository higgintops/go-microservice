package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// http handlers
	// HandleFunc is a convenience method
	// Registers a *function* to a path on the DefaultServeMux
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Println("hello world")
		body := req.Body

		// ioutil.ReadAll takes in an 'io.Reader', which means we can pass in anything
		// that implements an io.Reader
		data, err := ioutil.ReadAll(body)

		if err != nil {
			// if i got an error reading the response, i want to send back
			// the status code of bad request
			http.Error(rw, "Ooooops", http.StatusBadRequest)
			return
		}

		log.Printf("data %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}
