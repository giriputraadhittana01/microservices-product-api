package main

import (
	"context"
	"log"
	"microservicesapi/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Fourth
	// Apa yang mau dibuat dinamis maka itu yang diinject
	// Depency Injection
	lproducts := log.New(os.Stdout, "products-api ", log.LstdFlags)
	productsHandler := handlers.NewProductsHandler(lproducts)

	r := mux.NewRouter()
	// r.HandleFunc("/api/v1/product", productsHandler.GetAllProduct).Methods(http.MethodGet)
	// r.HandleFunc("/api/v1/product", productsHandler.AddProduct).Methods(http.MethodPost).Subrouter()
	// r.HandleFunc("/api/v1/product/{id}", productsHandler.UpdateProducts).Methods(http.MethodPut)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/product", productsHandler.GetAllProduct)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/v1/product", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareValidateProduct)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/v1/product/{id:[0-9]+}", productsHandler.UpdateProducts)
	putRouter.Use(productsHandler.MiddlewareValidateProduct)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
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
