package inventory

import (
	"reflect"
	"testing"
)

func TestInMemoryProductRepository_Get(t *testing.T) {
	type fields struct {
		products map[string]Product
	}
	type args struct {
		barcode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Product
		wantErr bool
	}{
		{
			name:    "Non initialised map",
			fields:  fields{products: nil},
			args:    args{barcode: "c0ffee"},
			want:    Product{},
			wantErr: true,
		},
		{
			name: "Existing product",
			fields: fields{map[string]Product{
				"c0ffee": {
					Barcode: "c0ffee",
				},
			}},
			args: args{barcode: "c0ffee"},
			want: Product{
				Barcode: "c0ffee",
			},
			wantErr: false,
		},
		{
			name: "Non existing product",
			fields: fields{map[string]Product{
				"c0ffee": {
					Barcode: "c0ffee",
				},
			}},
			args:    args{barcode: "m1lk"},
			want:    Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryProductRepository{
				products: tt.fields.products,
			}
			got, err := r.Get(tt.args.barcode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemoryProductRepository(t *testing.T) {
	r := NewInMemoryProductRepository()

	if r.products == nil {
		t.Error("expected repository to be initialised")
	}
}
