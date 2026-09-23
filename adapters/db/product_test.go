package adapters_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductDB_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(sqlmock.Sqlmock, string) (domain.IProduct, error)
	}{{
		name: "return error not found if db doesn't return a row",
		setupMock: func(mock sqlmock.Sqlmock, id string) (domain.IProduct, error) {
			mock.
				ExpectQuery(regexp.QuoteMeta("SELECT id, name, status, price FROM products WHERE id=$1")).
				WithArgs(id).
				WillReturnError(sql.ErrNoRows)

			return nil, errors.New("product not found")
		},
	}, {
		name: "return error if db return a different error",
		setupMock: func(mock sqlmock.Sqlmock, id string) (domain.IProduct, error) {
			mock.
				ExpectQuery(regexp.QuoteMeta("SELECT id, name, status, price FROM products WHERE id=$1")).
				WithArgs(id).
				WillReturnError(errors.New("database error"))

			return nil, errors.New("database error")
		},
	}, {
		name: "return product with success",
		setupMock: func(mock sqlmock.Sqlmock, id string) (domain.IProduct, error) {
			product := domain.NewProduct("any name", 0)
			product.ID = id
			rows := sqlmock.NewRows([]string{"id", "name", "status", "price"}).
				AddRow(product.ID, product.Name, product.Status, product.Price)
			mock.
				ExpectQuery(regexp.QuoteMeta("SELECT id, name, status, price FROM products WHERE id=$1")).
				WithArgs(id).
				WillReturnRows(rows)

			return product, nil
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			id := uuid.New().String()
			returnedProduct, expectedErr := tt.setupMock(mock, id)
			repository := NewProductDBSQLite(db)

			product, err := repository.Get(id)
			if expectedErr != nil {
				assert.Equal(t, expectedErr, err)
				assert.Nil(t, product)
				return
			}

			assert.Equal(t, returnedProduct, product)
			assert.Nil(t, err)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
