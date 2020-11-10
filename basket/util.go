package basket

import "github.com/google/uuid"

func generateID() string {
	return uuid.New().String()
}

func basketMapToSlice(baskets map[string]Basket) []Basket {
	s := make([]Basket, len(baskets))
	i := 0
	for _, val := range baskets {
		s[i] = val
		i++
	}

	return s
}
