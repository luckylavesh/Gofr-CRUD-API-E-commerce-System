package store

import (
	"database/sql"
	"fmt"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"sample/model"
)

type product struct{}

func New() *product {
	return &product{}
}

// Create inserts a new product record into the database
func (s *product) Create(ctx *gofr.Context, product *model.Product) (*model.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, createQuery, product.ID, product.Name, product.Price, product.Category)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return product, nil
}

// GetByID retrieves a product record based on its ID
func (s *product) GetByID(ctx *gofr.Context, id int) (*model.Product, error) {
	var resp model.Product

	err := ctx.DB().QueryRowContext(ctx, getByIDQuery, id).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Category)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.EntityNotFound{Entity: "product", ID: fmt.Sprintf("%v", id)}
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// Update updates an existing product record with the provided information
func (s *product) Update(ctx *gofr.Context, product *model.Product) (*model.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, updateQuery, product.Name, product.Price, product.Category, product.ID)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return product, nil
}

// Delete removes a product record from the database based on its ID
func (s *product) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
