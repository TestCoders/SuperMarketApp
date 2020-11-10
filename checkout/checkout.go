package checkout

type Purchasable interface {
	Total() int
	TotalExclDiscount() int
	TotalDiscount() int
}

// Order represent the checkout order
type Order struct {
	Number            int     `json:"number"`
	TotalExclDiscount float64 `json:"total_excl_discount"`
	TotalDiscount     float64 `json:"total_discount"`
	Total             float64 `json:"total"`
}

// MakePurchase converts the purchasable into an Order
func MakePurchase(purchasable Purchasable) Order {
	totalExclDiscount := float64(purchasable.TotalExclDiscount()) / 100
	totalDiscount := float64(purchasable.TotalDiscount()) / 100
	total := float64(purchasable.Total()) / 100

	return Order{
		TotalExclDiscount: totalExclDiscount,
		TotalDiscount:     totalDiscount,
		Total:             total,
	}
}
