package main

import (
	"context"
	"log"
	"microservicesapi/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Second
	// Apa yang mau dibuat dinamis maka itu yang diinject
	// Depency Injection
	lproducts := log.New(os.Stdout, "products-api ", log.LstdFlags)
	productsHandler := handlers.NewProductsHandler(lproducts)

	server := http.NewServeMux()

	server.Handle("/", productsHandler)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      server,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	log.Println("Receive Teriminate, Graceful Shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
