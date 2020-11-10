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
