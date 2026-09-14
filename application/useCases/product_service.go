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

func (ps *ProductServiceImpl) Get(id string) (domain.IProduct, error) {
	product, err := ps.persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductServiceImpl) Create(name string, price float32) (domain.IProduct, error) {
	product := domain.NewProduct(name, price)
	if _, err := product.IsValid(); err != nil {
		return nil, err
	}

	if err := ps.persistence.Save(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductServiceImpl) Enable(product domain.IProduct) error {
	if err := product.Enable(); err != nil {
		return err
	}

	if err := ps.persistence.Save(product); err != nil {
		return err
	}

	return nil
}

func (ps *ProductServiceImpl) Disable(product domain.IProduct) error {
	if err := product.Disable(); err != nil {
		return err
	}

	if err := ps.persistence.Save(product); err != nil {
		return err
	}

	return nil
}
