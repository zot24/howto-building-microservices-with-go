package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.00,
		SKU:   "abs-asd-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
