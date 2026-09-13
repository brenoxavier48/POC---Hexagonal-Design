package usecase_test

import (
	"errors"
	"testing"

	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/mocks"
	usecase "github.com/brenoxavier48/POC---Hexagonal-Design/application/useCases"
	"github.com/stretchr/testify/assert"
)

func TestProductService_Get(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(*mocks.MockProductPersistence) (domain.IProduct, error)
	}{{
		name: "return err if repository returns error",
		id:   "1234",
		setupMock: func(persistence *mocks.MockProductPersistence) (domain.IProduct, error) {
			persistence.
				On("Get", "1234").
				Return(nil, errors.New("Not found"))

			return nil, errors.New("Not found")
		},
	}, {
		name: "return product if it's been found",
		id:   "1234",
		setupMock: func(persistence *mocks.MockProductPersistence) (domain.IProduct, error) {
			product := &domain.Product{}
			persistence.
				On("Get", "1234").
				Return(nil, nil)

			return product, nil
		},
	}}

	for _, tt := range tests {
		persistence := new(mocks.MockProductPersistence)
		defer persistence.AssertNumberOfCalls(t, "Get", 1)
		expectedProduct, expectedError := tt.setupMock(persistence)

		service := usecase.NewProductService(persistence)

		product, err := service.Get(tt.id)
		if expectedError != nil {
			assert.Equal(t, expectedError, err)
			assert.Nil(t, product)
			return
		}

		assert.Equal(t, expectedProduct, product)
	}
}
