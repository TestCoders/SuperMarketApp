package basket

import (
	"github.com/marcelblijleven/SuperMarketApp/discount"
	"math"
)

// ProductLine represents a product line in the basket
type ProductLine struct {
	Barcode  string             `json:"barcode"`
	Name     string             `json:"name"`
	Quantity int                `json:"quantity"`
	Price    int                `json:"price"`
	Discount *discount.Discount `json:"discount,omitempty"`
}

// TotalDiscount returns the total discount of the product line
func (p *ProductLine) TotalDiscount() int {
	if p.Discount == nil {
		return 0
	}

	d := float64(p.Price*p.Quantity) / 100 * float64(p.Discount.Percentage)
	return int(math.Round(d))
}

// TotalExclDiscount returns the product line total excluding discount
func (p *ProductLine) TotalExclDiscount() int {
	return p.Quantity * p.Price
}

// Total returns the total of the product line
func (p *ProductLine) Total() int {
	return p.TotalExclDiscount() - p.TotalDiscount()
}

// NewProductLine returns a new product line with the provided values
func NewProductLine(barcode, name string, quantity, price int, discount *discount.Discount) ProductLine {
	return ProductLine{
		Barcode:  barcode,
		Name:     name,
		Quantity: quantity,
		Price:    price,
		Discount: discount,
	}
}
