package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productapi/data"
	"productapi/handlers"
	"time"

	protos "currentcyproject/protos/currency"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	// Eight
	// Apa yang mau dibuat dinamis maka itu yang diinject
	// Depency Injection
	lproducts := hclog.Default()
	vproducts := data.NewValidation()

	// Create Client
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cc := protos.NewCurrencyClient(conn)
	db := data.NewProductsDB(cc, lproducts)

	productsHandler := handlers.NewProductsHandler(lproducts, vproducts, db)

	r := mux.NewRouter()
	// r.HandleFunc("/api/v1/product", productsHandler.GetAllProduct).Methods(http.MethodGet)
	// r.HandleFunc("/api/v1/product", productsHandler.AddProduct).Methods(http.MethodPost).Subrouter()
	// r.HandleFunc("/api/v1/product/{id}", productsHandler.UpdateProducts).Methods(http.MethodPut)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/product", productsHandler.GetAllProduct).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/api/v1/product", productsHandler.GetAllProduct)
	getRouter.HandleFunc("/api/v1/product/{id:[0-9]+}", productsHandler.GetProduct).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/api/v1/product/{id:[0-9]+}", productsHandler.GetProduct)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/v1/product", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareValidateProduct)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/v1/product/{id:[0-9]+}", productsHandler.UpdateProducts)
	putRouter.Use(productsHandler.MiddlewareValidateProduct)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Private API
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	// Public API
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         ":8000",
		Handler:      ch(r),
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
