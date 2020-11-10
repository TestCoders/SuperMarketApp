package checkout

import (
	"testing"
)

type mockPurchasable struct {
}

func (m *mockPurchasable) Total() int {
	return 100
}

func (m *mockPurchasable) TotalExclDiscount() int {
	return 120
}

func (m *mockPurchasable) TotalDiscount() int {
	return 20
}

func TestMakePurchase(t *testing.T) {
	o := MakePurchase(&mockPurchasable{})

	if o.TotalDiscount != 0.20 {
		t.Error("total discount did not match expected value")
	}

	if o.TotalExclDiscount != 1.20 {
		t.Error("total excl discount did not match expected value")
	}

	if o.Total != 1.00 {
		t.Error("total did not match expected value")
	}
}
