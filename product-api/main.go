package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productapi/handlers"
	"time"
)

func main() {

	// create reference to handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(l)

	// servemux (aka a router)
	sm := http.NewServeMux()
	sm.Handle("/", productHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// ListenAndServe() normally blocks
	// so wrap in goroutine
	go func() {
		fmt.Println("starting server...")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	// Broadcast Kill/Interrupt is recieved
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until there IS a message to consume
	sig := <-sigChan
	l.Println("recieved terminate, graceful shutdown", sig)
	// Wait until requests complete, then shutdown
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

func NewHello(l *log.Logger) {
	panic("unimplemented")
}
