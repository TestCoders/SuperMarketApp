package basket

import (
	"reflect"
	"testing"
)

func TestInMemoryBasketRepository_Create(t *testing.T) {
	r := NewInMemoryBasketRepository()
	basket, _ := r.Create()

	if len(basket.ID) == 0 {
		t.Error("expected basket ID to be non nil")
	}

	retrievedBasket, _ := r.Get(basket.ID)
	if len(retrievedBasket.ID) == 0 {
		t.Error("expected basket with ID to be stored in map")
	}
}

func TestInMemoryBasketRepository_Delete(t *testing.T) {
	type fields struct {
		baskets map[string]Basket
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "Non existing basket",
			fields:  fields{baskets: map[string]Basket{}},
			args:    args{id: "abc"},
			wantErr: true,
			err:     ErrBasketNotFound,
		},
		{
			name: "Existing basket",
			fields: fields{baskets: map[string]Basket{
				"abc": {
					ID:           "abc",
					ProductLines: nil,
				},
			}},
			args:    args{id: "abc"},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InMemoryBasketRepository{
				baskets: tt.fields.baskets,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryBasketRepository_Get(t *testing.T) {
	type fields struct {
		baskets map[string]Basket
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Basket
		wantErr bool
	}{
		{
			name:    "Non existing basket id",
			fields:  fields{map[string]Basket{}},
			args:    args{id: "c0ffee"},
			want:    Basket{},
			wantErr: true,
		},
		{
			name: "Existing basket id",
			fields: fields{map[string]Basket{
				"c0ffee": {
					ID:           "c0ffee",
					ProductLines: nil,
				},
			}},
			args: args{id: "c0ffee"},
			want: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			},
			wantErr: false,
		},
		{
			name: "Multiple baskets",
			fields: fields{map[string]Basket{
				"c0ffee": {
					ID:           "c0ffee",
					ProductLines: nil,
				},
				"1337": {
					ID:           "1337",
					ProductLines: nil,
				},
			}},
			args: args{id: "c0ffee"},
			want: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryBasketRepository{
				baskets: tt.fields.baskets,
			}
			got, err := r.Get(tt.args.id)
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

func TestInMemoryBasketRepository_Update(t *testing.T) {
	type fields struct {
		baskets map[string]Basket
	}
	type args struct {
		basket Basket
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Basket
		wantErr bool
	}{
		{
			name:   "Non existing basket",
			fields: fields{baskets: map[string]Basket{}},
			args: args{basket: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			}},
			want: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			},
			wantErr: false,
		},
		{
			name: "Existing basket",
			fields: fields{baskets: map[string]Basket{
				"c0ffee": {
					ID:           "c0ffee",
					ProductLines: nil,
				},
			}},
			args: args{basket: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			}},
			want: Basket{
				ID:           "c0ffee",
				ProductLines: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InMemoryBasketRepository{
				baskets: tt.fields.baskets,
			}
			got, err := r.Update(tt.args.basket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemoryBasketRepository(t *testing.T) {
	r := NewInMemoryBasketRepository()

	if r.baskets == nil {
		t.Error("NewInMemoryBasketRepository().baskets = nil, want non nil")
	}
}
