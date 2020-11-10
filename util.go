package main

import (
	"github.com/marcelblijleven/SuperMarketApp/basket"
	"github.com/marcelblijleven/SuperMarketApp/inventory"
)

func productToProductLine(product inventory.Product, quantity int) basket.ProductLine {
	return basket.ProductLine{
		Barcode:  product.Barcode,
		Name:     product.Name,
		Quantity: quantity,
		Price:    product.Price,
		Discount: product.ActiveDiscount,
	}
}
