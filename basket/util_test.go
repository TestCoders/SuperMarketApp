package basket

import (
	"regexp"
	"testing"
)

func Test_generateID(t *testing.T) {
	r, err := regexp.Compile("^[0-9a-fA-F]{8}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{12}$")

	if err != nil {
		t.Fatal("error setting up regex")
		return
	}

	id := generateID()
	if m := r.MatchString(id); !m {
		t.Error("expected id to match uuid regex")
	}
}

func Test_basketMapToSlice(t *testing.T) {
	basketOne := Basket{
		ID:           "1",
		ProductLines: []ProductLine{},
	}
	basketTwo := Basket{
		ID:           "2",
		ProductLines: []ProductLine{},
	}

	m := map[string]Basket{
		"1": basketOne,
		"2": basketTwo,
	}

	s := basketMapToSlice(m)

	if len(s) != 2 {
		t.Error("expected len of slice to be 2")
		return
	}

	if s[0].ID != "1" {
		t.Errorf("expected id of basket to match 1, instead got %q", s[0].ID)
	}

	if s[0].ProductLines == nil {
		t.Error("expected ProductLines to be non nil")
	}

	if s[1].ID != "2" {
		t.Errorf("expected id of basket to match 2, instead got %q", s[0].ID)
	}

	if s[1].ProductLines == nil {
		t.Error("expected ProductLines to be non nil")
	}
}
