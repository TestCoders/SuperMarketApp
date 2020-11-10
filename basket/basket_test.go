package basket

import (
	"github.com/marcelblijleven/SuperMarketApp/discount"
	"testing"
)

var (
	testDiscountOne = &discount.Discount{
		Description: "Test Discount One",
		Percentage:  20,
	}
	testDiscountTwo = &discount.Discount{
		Description: "Test Discount Two",
		Percentage:  35,
	}
)

var (
	testProductLineOne   = NewProductLine("123", "Test Product Line One", 1, 1500, nil)
	testProductLineTwo   = NewProductLine("123", "Test Product Line Two", 2, 1500, nil)
	testProductLineThree = NewProductLine("123", "Test Product Line Three", 1, 1500, testDiscountOne)
	testProductLineFour  = NewProductLine("123", "Test Product Line Four", 2, 1500, testDiscountTwo)
)

func TestBasket_Total(t *testing.T) {
	type fields struct {
		ID           string
		ProductLines []ProductLine
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Two products without discounts",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineTwo,
			}},
			want: 4500,
		},
		{
			name: "One product without, one product with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineThree,
			}},
			want: 2700,
		},
		{
			name: "Two products with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineThree,
				testProductLineFour,
			}},
			want: 3150,
		},
		{
			name: "Two products without, two products with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineTwo,
				testProductLineThree,
				testProductLineFour,
			}},
			want: 7650,
		},
		{
			name:   "No products in basket",
			fields: fields{ProductLines: []ProductLine{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Basket{
				ID:           tt.fields.ID,
				ProductLines: tt.fields.ProductLines,
			}
			if got := b.Total(); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasket_TotalDiscount(t *testing.T) {
	type fields struct {
		Products []ProductLine
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Single product with discount",
			fields: fields{
				Products: []ProductLine{testProductLineThree},
			},
			want: 300,
		},
		{
			name: "Two products with discount",
			fields: fields{
				Products: []ProductLine{
					testProductLineThree,
					testProductLineFour,
				},
			},
			want: 1350,
		},
		{
			name: "No products with discount",
			fields: fields{
				Products: []ProductLine{
					testProductLineOne,
				},
			},
			want: 0,
		},
		{
			name:   "No products in basket",
			fields: fields{Products: []ProductLine{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Basket{
				ProductLines: tt.fields.Products,
			}
			if got := b.TotalDiscount(); got != tt.want {
				t.Errorf("TotalDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasket_TotalExclDiscount(t *testing.T) {
	type fields struct {
		ID           string
		ProductLines []ProductLine
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Two products without discounts",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineTwo,
			}},
			want: 4500,
		},
		{
			name: "One product without, one product with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineThree,
			}},
			want: 3000,
		},
		{
			name: "Two products with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineThree,
				testProductLineFour,
			}},
			want: 4500,
		},
		{
			name: "Two products without, two products with discount",
			fields: fields{ProductLines: []ProductLine{
				testProductLineOne,
				testProductLineTwo,
				testProductLineThree,
				testProductLineFour,
			}},
			want: 9000,
		},
		{
			name:   "No products in basket",
			fields: fields{ProductLines: []ProductLine{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Basket{
				ID:           tt.fields.ID,
				ProductLines: tt.fields.ProductLines,
			}
			if got := b.TotalExclDiscount(); got != tt.want {
				t.Errorf("TotalExclDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasket_UpdateProductLine(t *testing.T) {
	type fields struct {
		ID           string
		ProductLines []ProductLine
	}
	type args struct {
		line ProductLine
	}
	type want struct {
		barcode  string
		quantity int
		total    int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Update product in empty basket",
			fields: fields{
				ID:           "abc123",
				ProductLines: []ProductLine{},
			},
			args: args{ProductLine{
				Barcode:  "1337",
				Quantity: 1,
				Price:    2000,
				Discount: nil,
			}},
			want: want{
				barcode:  "1337",
				quantity: 1,
				total:    1,
			},
		},
		{
			name: "Update existing product",
			fields: fields{
				ID: "abc123",
				ProductLines: []ProductLine{
					{
						Barcode:  "1337",
						Quantity: 1,
						Price:    2000,
						Discount: nil,
					},
				},
			},
			args: args{ProductLine{
				Barcode:  "1337",
				Quantity: 1,
				Price:    2000,
				Discount: nil,
			}},
			want: want{
				barcode:  "1337",
				quantity: 2,
				total:    2,
			},
		},
		{
			name: "Update nonexisting product",
			fields: fields{
				ID: "abc123",
				ProductLines: []ProductLine{
					{
						Barcode:  "1337",
						Quantity: 1,
						Price:    2000,
						Discount: nil,
					},
				},
			},
			args: args{ProductLine{
				Barcode:  "c0ffee",
				Quantity: 1,
				Price:    1500,
				Discount: nil,
			}},
			want: want{
				barcode:  "c0ffee",
				quantity: 1,
				total:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Basket{
				ID:           tt.fields.ID,
				ProductLines: tt.fields.ProductLines,
			}

			b.UpdateProductLine(tt.args.line)

			totalQuantity := 0

			for _, p := range b.ProductLines {
				if p.Barcode == tt.want.barcode {
					if p.Quantity != tt.want.quantity {
						t.Errorf("UpdateProduct() = %v, want %v", p.Quantity, tt.want.quantity)
					}
				}

				totalQuantity += p.Quantity
			}

			if totalQuantity != tt.want.total {
				t.Errorf("UpdateProduct() total = %v, want %v", totalQuantity, tt.want.total)
			}
		})
	}
}
