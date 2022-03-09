package handlers

import (
	"net/http"
	"productapi/data"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type KeyProduct struct{}

type Products struct {
	l         hclog.Logger
	v         *data.Validation
	productDB *data.ProductsDB
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewProductsHandler(l hclog.Logger, v *data.Validation, pdb *data.ProductsDB) *Products {
	return &Products{l, v, pdb}
}

// swagger:route GET /api/v1/product products GetAllProduct
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetAllProduct handles GET requests and returns all current products
func (p *Products) GetAllProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	cur := r.URL.Query().Get("currency")
	listOfProduct, err := p.productDB.GetProducts(cur)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	err = data.ToJSON(listOfProduct, rw)
	// json.NewEncoder(rw).Encode(listOfProduct)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

// swagger:route GET /api/v1/product/{id} products GetProduct
// Return a product from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetProduct handles GET requests
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	cur := r.URL.Query().Get("currency")
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product, err := p.productDB.GetProduct(id, cur)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	err = data.ToJSON(product, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

// swagger:route POST /product products CreateProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.productDB.AddProduct(&prod)
}

// swagger:route PUT /api/v1/product/{id} products UpdateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = p.productDB.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
