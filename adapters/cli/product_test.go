package adapters_test

import (
	"errors"
	"fmt"
	"testing"

	adapters "github.com/brenoxavier48/POC---Hexagonal-Design/adapters/cli"
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductCLI(t *testing.T) {
	tests := []struct {
		name              string
		setupMock         func(*mocks.MockProductService) (domain.IProduct, error)
		getExpectedResult func(domain.IProduct) string
		action            string
		productId         string
		productName       string
		price             float32
	}{{
		name: "return error if service return error at create",
		setupMock: func(service *mocks.MockProductService) (domain.IProduct, error) {
			expectedError := errors.New("service error")
			service.
				On("Create", "any name", float32(2.0)).
				Return(nil, expectedError)
			return nil, expectedError
		},
		getExpectedResult: func(product domain.IProduct) string { return "" },
		action:            "create",
		productId:         "",
		productName:       "any name",
		price:             float32(2.0),
	}, {
		name: "return valid result and nil error",
		setupMock: func(service *mocks.MockProductService) (domain.IProduct, error) {
			expectedProduct := domain.NewProduct("any name", float32(2.0))
			service.
				On("Create", expectedProduct.Name, expectedProduct.Price).
				Return(expectedProduct, nil)
			return expectedProduct, nil
		},
		getExpectedResult: func(product domain.IProduct) string {
			return fmt.Sprintf("Product with the name %s has been created with the price %f and status %s",
				product.GetName(), product.GetPrice(), product.GetStatus())
		},
		action:      "create",
		productId:   "",
		productName: "any name",
		price:       float32(2.0),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := new(mocks.MockProductService)
			expectedProduct, expectedErr := tt.setupMock(service)

			result, err := adapters.Run(service, tt.action, tt.productId, tt.productName, tt.price)

			if expectedErr != nil {
				assert.Equal(t, expectedErr, err)
				assert.Equal(t, "", result)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.getExpectedResult(expectedProduct), result)
		})
	}
}
