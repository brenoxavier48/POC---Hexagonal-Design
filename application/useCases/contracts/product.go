package usecase

import (
	domain "github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
)

type IProductService interface {
	Get(id string) (domain.IProduct, error)
	Create(name string, price float32) (domain.IProduct, error)
	Enable(product domain.IProduct) error
	Disable(product domain.IProduct) error
}

type IProductPersistence interface {
	Get(id string) (domain.IProduct, error)
	Save(product domain.IProduct) error
}
