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

type MockProductService struct {
	mock.Mock
}

func (mPS *MockProductService) Get(id string) (domain.IProduct, error) {
	args := mPS.Called(id)

	product, _ := args.Get(0).(domain.IProduct)

	return product, args.Error(1)
}

func (mPS *MockProductService) Create(name string, price float32) (domain.IProduct, error) {
	args := mPS.Called(name, price)

	product, _ := args.Get(0).(domain.IProduct)

	return product, args.Error(1)
}

func (mPS *MockProductService) Disable(product domain.IProduct) error {
	args := mPS.Called(product)
	return args.Error(0)
}

func (mPS *MockProductService) Enable(product domain.IProduct) error {
	args := mPS.Called(product)
	return args.Error(0)
}
