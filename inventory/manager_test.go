package inventory

import (
	"testing"
)

type mockRepository struct {
}

func (m *mockRepository) Get(barcode string) (Product, error) {
	return Product{Barcode: barcode, Name: "Get called"}, nil
}

func TestManager_Get(t *testing.T) {
	m := NewManager(&mockRepository{})
	p, _ := m.Get("c0ffee")

	if p.Name != "Get called" {
		t.Error("incorrect repository called")
	}
}

func TestNewManager(t *testing.T) {
	m := NewManager(&mockRepository{})

	if m.repository == nil {
		t.Error("expected repository to be initialised")
	}
}
