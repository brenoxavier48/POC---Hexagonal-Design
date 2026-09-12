package application

import "github.com/google/uuid"

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
	ID     uuid.UUID
	Name   string
	Status string
	Price  float32
}

func (p *Product) IsValid() (bool, error) {
	return false, nil
}

func (p *Product) Enable() error {
	return nil
}

func (p *Product) Disable() error {
	return nil
}
