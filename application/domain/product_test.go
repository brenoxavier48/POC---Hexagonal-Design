package domain_test

import (
	"errors"

	domain "github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestProduct_Enable(t *testing.T) {
	tests := []struct {
		name        string
		product     domain.Product
		expectedErr error
	}{{
		name:        "return error if price is zero",
		product:     domain.Product{Price: 0},
		expectedErr: errors.New("ENABLE ERROR: price should be greater then zero"),
	}, {
		name:        "enable product if price is greater then zero",
		product:     domain.Product{Price: 1},
		expectedErr: nil,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Enable()

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				assert.Equal(t, domain.DISABLED, tt.product.Status)
				return
			}

			assert.Equal(t, domain.ENABLED, tt.product.Status)
		})
	}
}

func TestProduct_Disable(t *testing.T) {
	tests := []struct {
		name        string
		product     domain.Product
		expectedErr error
	}{{
		name:        "return error if product has a price greater then zero",
		product:     domain.Product{Price: 1},
		expectedErr: errors.New("DISABLE ERROR: price should be zero"),
	}, {
		name:        "return nil if product has price equals to zero",
		product:     domain.Product{Price: 0},
		expectedErr: nil,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Disable()

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			require.Nil(t, err)
			assert.Equal(t, domain.DISABLED, tt.product.Status)
		})
	}
}

func TestProduct_IsValid(t *testing.T) {
	tests := []struct {
		name        string
		product     domain.Product
		expectedErr string
	}{{
		name:        "return error if status is empty",
		product:     domain.Product{Status: "invalid status"},
		expectedErr: "product has to have a valid status",
	}, {
		name:        "return error if status is different then enable or disable",
		product:     domain.Product{Status: "invalid status"},
		expectedErr: "product has to have a valid status",
	}, {
		name:        "return error if price is lower then zero",
		product:     domain.Product{Status: domain.ENABLED, Price: -1},
		expectedErr: "price is lower then zero",
	}, {
		name:        "return error if id is not valid",
		product:     domain.Product{Status: domain.ENABLED, Price: 0, Name: "p1"},
		expectedErr: "ID: non zero value required",
	}, {
		name:        "return error if name is missing",
		product:     domain.Product{Status: domain.ENABLED, Price: 0, ID: uuid.New().String()},
		expectedErr: "Name: non zero value required",
	}, {
		name:        "return true if product is validated",
		product:     domain.Product{Status: domain.ENABLED, Price: 0, ID: uuid.New().String(), Name: "p1"},
		expectedErr: "",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, err := tt.product.IsValid()

			if tt.expectedErr != "" {
				assert.False(t, isValid)
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}

			assert.True(t, isValid)
			assert.Nil(t, err)
		})
	}
}
