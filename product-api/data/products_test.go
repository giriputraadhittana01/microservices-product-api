package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:        "nics",
		Description: "Hello World",
		Price:       1.00,
		SKU:         "abc-def-ghi",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
