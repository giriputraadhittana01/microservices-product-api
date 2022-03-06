package handlers

import (
	"encoding/json"
	"log"
	"microservicesapi/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.GetAllProduct(rw, r)
		return
	}

	if r.Method == http.MethodGet {
		p.AddProduct(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetAllProduct(rw http.ResponseWriter, r *http.Request) {
	listOfProduct := data.GetProducts()
	// _ = listOfProduct.ToJSON(rw)
	json.NewEncoder(rw).Encode(listOfProduct)
}
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

}
