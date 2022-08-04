package handlers

import (
	"fmt"
	"log"
	"net/http"
	"productapi/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// This signature is what we need to satisfy the http.Handler interface:
// see here: https://pkg.go.dev/net/http#Handler which defines a method
// ServeHTTP with two inputs a ResponseWriter and a *Request
func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Made it inside the handler for products")

	listProducts := data.GetProducts()

	// convert listProducts into JSON
	err := listProducts.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Could not Marshall json products", http.StatusInternalServerError)
		return
	}
}
