package mocks

import (
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductPersistence struct {
	mock.Mock
}

func (m *MockProductPersistence) Get(id string) (domain.IProduct, error) {
	args := m.Called(id)

	product, _ := args.Get(0).(domain.IProduct)

	return product, args.Error(1)
}

func (m *MockProductPersistence) Save(product domain.IProduct) error {
	args := m.Called(product)
	return args.Error(0)
}
