package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:        "nics",
		Description: "Hello World",
		Price:       1.00,
		SKU:         "abc-def-ghi",
	}
	vproducts := NewValidation()
	err := vproducts.Validate(p)
	if len(err) != 0 {
		t.Fatal(err)
	}
}
