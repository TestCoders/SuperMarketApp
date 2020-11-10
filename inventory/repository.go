package inventory

import "github.com/marcelblijleven/SuperMarketApp/discount"

type Repository interface {
	Get(barcode string) (Product, error)
}

type InMemoryProductRepository struct {
	products map[string]Product
}

func (r *InMemoryProductRepository) Get(barcode string) (Product, error) {
	if r.products == nil {
		return Product{}, ErrRepositoryNotInitialised
	}

	p, ok := r.products[barcode]

	if !ok {
		return Product{}, ErrProductNotFound
	}

	return p, nil
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	r := &InMemoryProductRepository{map[string]Product{
		"156734": {
			Barcode:        "156734",
			Name:           "Kaas",
			Price:          499,
			ActiveDiscount: nil,
		},
		"579843": {
			Barcode:        "579843",
			Name:           "Ham",
			Price:          149,
			ActiveDiscount: discount.BonusDiscount,
		},
		"378941": {
			Barcode:        "378941",
			Name:           "Melk",
			Price:          99,
			ActiveDiscount: discount.OutOfDateDiscount,
		},
		"739214": {
			Barcode:        "739214",
			Name:           "Pizza",
			Price:          459,
			ActiveDiscount: nil,
		},
		"798234": {
			Barcode:        "798234",
			Name:           "Bier",
			Price:          1199,
			ActiveDiscount: discount.BonusDiscount,
		},
	}}
	return r
}
