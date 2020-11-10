package basket

import (
	"github.com/marcelblijleven/SuperMarketApp/discount"
	"testing"
)

func TestNewProductLine(t *testing.T) {
	code := "c0ffee"
	name := "Koffie"
	quantity := 2
	price := 1234
	d := testDiscountOne

	p := NewProductLine(code, name, quantity, price, d)

	if code != p.Barcode || name != p.Name || quantity != p.Quantity || price != p.Price || p.Discount != d {
		t.Error("expected ProductLine to have correct values")
	}
}

func TestProductLine_Total(t *testing.T) {
	type fields struct {
		Barcode  string
		Name     string
		Quantity int
		Price    int
		Discount *discount.Discount
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Line without discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: nil,
			},
			want: 1499,
		},
		{
			name: "Line with discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 1199,
		},
		{
			name: "Line without discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: nil,
			},
			want: 2998,
		},
		{
			name: "Line with discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 2398,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductLine{
				Barcode:  tt.fields.Barcode,
				Name:     tt.fields.Name,
				Quantity: tt.fields.Quantity,
				Price:    tt.fields.Price,
				Discount: tt.fields.Discount,
			}
			if got := p.Total(); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductLine_TotalDiscount(t *testing.T) {
	type fields struct {
		Barcode  string
		Name     string
		Quantity int
		Price    int
		Discount *discount.Discount
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Line without discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: nil,
			},
			want: 0,
		},
		{
			name: "Line with discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 300,
		},
		{
			name: "Line without discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: nil,
			},
			want: 0,
		},
		{
			name: "Line with discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductLine{
				Barcode:  tt.fields.Barcode,
				Name:     tt.fields.Name,
				Quantity: tt.fields.Quantity,
				Price:    tt.fields.Price,
				Discount: tt.fields.Discount,
			}
			if got := p.TotalDiscount(); got != tt.want {
				t.Errorf("TotalDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductLine_TotalExclDiscount(t *testing.T) {
	type fields struct {
		Barcode  string
		Name     string
		Quantity int
		Price    int
		Discount *discount.Discount
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Line without discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: nil,
			},
			want: 1499,
		},
		{
			name: "Line with discount",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 1,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 1499,
		},
		{
			name: "Line without discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: nil,
			},
			want: 2998,
		},
		{
			name: "Line with discount and quantity of 2",
			fields: fields{
				Barcode:  "c0ffee",
				Name:     "Koffie",
				Quantity: 2,
				Price:    1499,
				Discount: testDiscountOne,
			},
			want: 2998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductLine{
				Barcode:  tt.fields.Barcode,
				Name:     tt.fields.Name,
				Quantity: tt.fields.Quantity,
				Price:    tt.fields.Price,
				Discount: tt.fields.Discount,
			}
			if got := p.TotalExclDiscount(); got != tt.want {
				t.Errorf("TotalExclDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}
