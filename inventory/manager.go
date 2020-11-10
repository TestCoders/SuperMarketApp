package inventory

type Manager struct {
	repository Repository
}

func (m *Manager) Get(barcode string) (Product, error) {
	return m.repository.Get(barcode)
}

func NewManager(repository Repository) Manager {
	return Manager{repository: repository}
}
