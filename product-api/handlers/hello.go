package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	fmt.Println("IN HERE")
	return &Hello{l}
}

// This signature is what we need to satisfy the http.Handler interface:
// see here: https://pkg.go.dev/net/http#Handler which defines a method
// ServeHTTP with two inputs a ResponseWriter and a *Request
func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("hello world")
	body := req.Body

	data, err := ioutil.ReadAll(body)

	if err != nil {
		http.Error(rw, "Ooooops", http.StatusBadRequest)
		return
	}

	fmt.Printf("data %s", data)
}
