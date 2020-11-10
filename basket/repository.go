package basket

// Repository interface
type Repository interface {
	GetAll() ([]Basket, error)
	Get(id string) (Basket, error)
	Create() (Basket, error)
	Save(basket Basket) error
	Update(basket Basket) (Basket, error)
	Delete(id string) error
}

// InMemoryBasketRepository is a basket repository which can be used
// for testing purposes or when data persistence between runs is not needed
type InMemoryBasketRepository struct {
	baskets map[string]Basket
}

// NewInMemoryBasketRepository returns a InMemoryBasketRepository with an initialised
// map[string]Basket
func NewInMemoryBasketRepository() *InMemoryBasketRepository {
	return &InMemoryBasketRepository{baskets: map[string]Basket{}}
}

func (r *InMemoryBasketRepository) Save(basket Basket) error {
	r.baskets[basket.ID] = basket
	return nil
}

func (r *InMemoryBasketRepository) GetAll() ([]Basket, error) {
	if r.baskets == nil {
		return nil, ErrRepositoryNotInitialised
	}

	return basketMapToSlice(r.baskets), nil
}

// Get retrieves a basket by id
func (r *InMemoryBasketRepository) Get(id string) (Basket, error) {
	if r.baskets == nil {
		return Basket{}, ErrRepositoryNotInitialised
	}

	basket, ok := r.baskets[id]

	if !ok {
		return basket, ErrBasketNotFound
	}

	return basket, nil
}

// Create creates a basket with a generated ID
func (r InMemoryBasketRepository) Create() (Basket, error) {
	if r.baskets == nil {
		return Basket{}, ErrRepositoryNotInitialised
	}

	basket := Basket{
		ID:           generateID(),
		ProductLines: []ProductLine{},
	}

	r.baskets[basket.ID] = basket
	return basket, nil
}

// Update replaces the basket in the map with the provided basket
func (r InMemoryBasketRepository) Update(basket Basket) (Basket, error) {
	if r.baskets == nil {
		return Basket{}, ErrRepositoryNotInitialised
	}

	r.baskets[basket.ID] = basket
	return basket, nil
}

// Delete removes the basket from the map by ID
func (r InMemoryBasketRepository) Delete(id string) error {
	if r.baskets == nil {
		return ErrRepositoryNotInitialised
	}

	_, ok := r.baskets[id]

	if ok {
		delete(r.baskets, id)
		return nil
	}

	return ErrBasketNotFound
}
