package adapters

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/brenoxavier48/POC---Hexagonal-Design/application/domain"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDBSQLite struct {
	db *sql.DB
}

func NewProductDBSQLite(db *sql.DB) *ProductDBSQLite {
	return &ProductDBSQLite{db}
}

func (pDB *ProductDBSQLite) Get(id string) (domain.IProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	person := &domain.Product{}

	err := pDB.db.QueryRowContext(ctx, "SELECT id, name, status, price FROM products WHERE id=$1", id).
		Scan(&person.ID, &person.Name, &person.Status, &person.Price)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	return person, nil
}

func (pDB *ProductDBSQLite) Save(product domain.IProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	productToPersist := domain.NewProduct(product.GetName(), product.GetPrice())
	productToPersist.Status = product.GetStatus()
	query := `
		INSERT INTO products (id, name, status, price) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT (id) 
		DO UPDATE SET 
			name = EXCLUDED.name
			status = EXCLUDED.status
			price = EXCLUDED.price
	`

	_, err := pDB.db.ExecContext(
		ctx,
		query,
		productToPersist.ID,
		productToPersist.Name,
		productToPersist.Status,
		productToPersist.Price,
	)
	if err != nil {
		return err
	}

	return nil
}
