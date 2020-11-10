package main

import (
	"github.com/marcelblijleven/SuperMarketApp/inventory"
	"testing"
)

func Test_productToProductLine(t *testing.T) {
	product := inventory.Product{
		Barcode:        "c0ffee",
		Name:           "Koffie",
		Price:          149,
		ActiveDiscount: nil,
	}

	line := productToProductLine(product, 2)

	if line.Quantity != 2 {
		t.Error("expected quantity to match")
	}

	if line.Barcode != product.Barcode {
		t.Error("expected barcode to match")
	}

	if line.Name != product.Name {
		t.Error("expected name to match")
	}

	if line.Price != product.Price {
		t.Error("expected price to match")
	}

	if line.Discount != product.ActiveDiscount {
		t.Error("expected discount to match")
	}
}
