package application_test

import (
	"errors"

	"github.com/brenoxavier48/POC---Hexagonal-Design/application"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestProduct_Enable(t *testing.T) {
	tests := []struct {
		name        string
		product     application.Product
		expectedErr error
	}{{
		name:        "return error if price is zero",
		product:     application.Product{Price: 0},
		expectedErr: errors.New("ENABLE ERROR: price should be greater then zero"),
	}, {
		name:        "enable product if price is greater then zero",
		product:     application.Product{Price: 1},
		expectedErr: nil,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Enable()

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				assert.Equal(t, application.DISABLED, tt.product.Status)
				return
			}

			assert.Equal(t, application.ENABLED, tt.product.Status)
		})
	}
}
