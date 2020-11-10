package basket

// Basket represents a shopping basket
type Basket struct {
	ID           string        `json:"barcode"`
	ProductLines []ProductLine `json:"products"`
}

// Total returns the basket total including discounts
func (b *Basket) Total() int {
	total := 0

	for _, product := range b.ProductLines {
		total += product.Total()
	}

	return total
}

// TotalExclDiscount returns the basket total excluding discounts
func (b *Basket) TotalExclDiscount() int {
	total := 0

	for _, product := range b.ProductLines {
		total += product.TotalExclDiscount()
	}

	return total
}

// TotalDiscount returns the total discount of the basket
func (b *Basket) TotalDiscount() int {
	total := 0

	for _, product := range b.ProductLines {
		total += product.TotalDiscount()
	}

	return total
}

// UpdateProductLine updates the quantity of an existing product line,
// or adds the product line to the basket if it didn't exist yet
func (b *Basket) UpdateProductLine(line ProductLine) {
	var found bool

	for i := range b.ProductLines {
		if b.ProductLines[i].Barcode == line.Barcode {
			found = true
			b.ProductLines[i].Quantity += line.Quantity
			return
		}
	}

	if !found {
		b.ProductLines = append(b.ProductLines, line)
	}
}
