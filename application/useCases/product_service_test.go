package usecase_test

import (
	"errors"
	"testing"

	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/mocks"
	usecase "github.com/brenoxavier48/POC---Hexagonal-Design/application/useCases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestProductService_Create(t *testing.T) {
	tests := []struct {
		name               string
		productName        string
		productPrice       float32
		setupMock          func(*mocks.MockProductPersistence) error
		expectedProductErr string
	}{{
		name:               "return err if product name is not valid",
		productName:        "",
		productPrice:       0,
		setupMock:          func(persistence *mocks.MockProductPersistence) error { return nil },
		expectedProductErr: "Name: non zero value required",
	}, {
		name:               "return err if product price is not valid",
		productName:        "p1",
		productPrice:       -1,
		setupMock:          func(persistence *mocks.MockProductPersistence) error { return nil },
		expectedProductErr: "price is lower then zero",
	}, {
		name:         "return err if persistence fails",
		productName:  "p1",
		productPrice: 0,
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			expectedErr := errors.New("database error")
			persistence.
				On("Save", mock.Anything).
				Return(expectedErr)
			return expectedErr
		},
		expectedProductErr: "",
	}, {
		name:         "return product if persistence works",
		productName:  "p1",
		productPrice: 0,
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			persistence.
				On("Save", mock.Anything).
				Return(nil)
			return nil
		},
		expectedProductErr: "",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			persistence := new(mocks.MockProductPersistence)
			expectedPersistenceErr := tt.setupMock(persistence)

			service := usecase.NewProductService(persistence)

			product, err := service.Create(tt.productName, tt.productPrice)

			if tt.expectedProductErr != "" {
				assert.Equal(t, tt.expectedProductErr, err.Error())
				persistence.AssertNotCalled(t, "Save")
				assert.Nil(t, product)
				return
			}

			defer persistence.AssertNumberOfCalls(t, "Save", 1)
			if expectedPersistenceErr != nil {
				assert.Equal(t, expectedPersistenceErr, err)
				assert.Nil(t, product)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, product)
		})
	}
}

func TestProductService_Enable(t *testing.T) {
	tests := []struct {
		name               string
		product            domain.IProduct
		setupMock          func(*mocks.MockProductPersistence) error
		expectedProductErr string
	}{{
		name:               "return err if product invalid to enable",
		product:            &domain.Product{},
		setupMock:          func(persistence *mocks.MockProductPersistence) error { return nil },
		expectedProductErr: "ENABLE ERROR: price should be greater then zero",
	}, {
		name:    "return err if database fails",
		product: &domain.Product{Price: 2.0},
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			expectedError := errors.New("database error")
			persistence.
				On("Save", mock.Anything).
				Return(expectedError)
			return expectedError
		},
		expectedProductErr: "",
	}, {
		name:    "enable product if persistence works",
		product: &domain.Product{Price: 2.0},
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			persistence.
				On("Save", mock.Anything).
				Return(nil)
			return nil
		},
		expectedProductErr: "",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			persistence := new(mocks.MockProductPersistence)
			expectedPersistenceErr := tt.setupMock(persistence)

			service := usecase.NewProductService(persistence)

			err := service.Enable(tt.product)

			if tt.expectedProductErr != "" {
				assert.Equal(t, tt.expectedProductErr, err.Error())
				persistence.AssertNotCalled(t, "Save")
				return
			}

			defer persistence.AssertNumberOfCalls(t, "Save", 1)
			if expectedPersistenceErr != nil {
				assert.Equal(t, expectedPersistenceErr, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestProductService_Disable(t *testing.T) {
	tests := []struct {
		name               string
		product            domain.IProduct
		setupMock          func(*mocks.MockProductPersistence) error
		expectedProductErr string
	}{{
		name:               "return err if product invalid to disable",
		product:            &domain.Product{Price: 2.0},
		setupMock:          func(persistence *mocks.MockProductPersistence) error { return nil },
		expectedProductErr: "DISABLE ERROR: price should be zero",
	}, {
		name:    "return err if database fails",
		product: &domain.Product{},
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			expectedError := errors.New("database error")
			persistence.
				On("Save", mock.Anything).
				Return(expectedError)
			return expectedError
		},
		expectedProductErr: "",
	}, {
		name:    "enable product if persistence works",
		product: &domain.Product{Price: 0},
		setupMock: func(persistence *mocks.MockProductPersistence) error {
			persistence.
				On("Save", mock.Anything).
				Return(nil)
			return nil
		},
		expectedProductErr: "",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			persistence := new(mocks.MockProductPersistence)
			expectedPersistenceErr := tt.setupMock(persistence)

			service := usecase.NewProductService(persistence)

			err := service.Disable(tt.product)

			if tt.expectedProductErr != "" {
				assert.Equal(t, tt.expectedProductErr, err.Error())
				persistence.AssertNotCalled(t, "Save")
				return
			}

			defer persistence.AssertNumberOfCalls(t, "Save", 1)
			if expectedPersistenceErr != nil {
				assert.Equal(t, expectedPersistenceErr, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
