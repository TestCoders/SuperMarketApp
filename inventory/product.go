package inventory

import "github.com/marcelblijleven/SuperMarketApp/discount"

type Product struct {
	Barcode        string             `json:"barcode"`
	Name           string             `json:"name"`
	Price          int                `json:"price"`
	ActiveDiscount *discount.Discount `json:"discount"`
}
