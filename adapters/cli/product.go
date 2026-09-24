package adapters

import (
	"fmt"

	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	usecase "github.com/brenoxavier48/POC---Hexagonal-Design/application/useCases/contracts"
)

func runCreate(service usecase.ProductService, productId string, productName string, price float32) (domain.IProduct, error) {
	product, err := service.Create(productName, price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func Run(service usecase.ProductService, action string, productId string, productName string, price float32) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product with the name %s has been created with the price %f and status %s",
			product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", product.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Name: %s\nPrice: %f\nStatus: %s",
			res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil
}
