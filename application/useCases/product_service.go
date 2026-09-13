package usecase

import (
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	usecase "github.com/brenoxavier48/POC---Hexagonal-Design/application/useCases/contracts"
)

type ProductServiceImpl struct {
	persistence usecase.ProductPersistence
}

func NewProductService(persistence usecase.ProductPersistence) *ProductServiceImpl {
	return &ProductServiceImpl{persistence}
}

func (ps *ProductServiceImpl) Get(id string) (*domain.IProduct, error) {
	product, err := ps.persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
