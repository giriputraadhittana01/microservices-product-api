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
	// First
	// Apa yang mau dibuat dinamis maka itu yang diinject
	// Depency Injection
	lhello := log.New(os.Stdout, "hello-api", log.LstdFlags)
	lgoodbye := log.New(os.Stdout, "goodbye-api", log.LstdFlags)
	helloHandler := handlers.NewHelloHandler(lhello)
	goodbyeHandler := handlers.NewGoodByeHandler(lgoodbye)
	// Tidak Depency Injection maka susah untuk diubah
	hyHandler := handlers.GetAllDataHy

	server := http.NewServeMux()

	server.HandleFunc("/hello", helloHandler.GetAllDataHello)
	server.HandleFunc("/goodbye", goodbyeHandler.GetAllDataGoodBye)
	server.HandleFunc("/hy", hyHandler)

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
