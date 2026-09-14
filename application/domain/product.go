package domain

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Product struct {
	ID     string  `valid:"required,uuid"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float32 `valid:"float32,optional"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		ID:     uuid.New().String(),
		Status: DISABLED,
		Name:   name,
		Price:  price,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("product has to have a valid status")
	}

	if p.Price < 0 {
		return false, errors.New("price is lower then zero")
	}

	if result, err := govalidator.ValidateStruct(p); err != nil {
		fmt.Println(result)
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		p.Status = DISABLED
		return errors.New("ENABLE ERROR: price should be greater then zero")
	}
	p.Status = ENABLED

	return nil
}

func (p *Product) Disable() error {
	if p.Price != 0 {
		return errors.New("DISABLE ERROR: price should be zero")
	}
	p.Status = DISABLED

	return nil
}
