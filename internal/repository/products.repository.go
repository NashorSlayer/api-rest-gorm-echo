package repository

import (
	"context"

	entity "github.com/rarifbkn/api-rest-gorm-echo/internal/entity"
)

const (
	qryInserProduct = `
		insert into PRODUCTS (name,description,price,created_by) values ($1,$2,$3,$4);
	`

	qryGetAllProducts = `
		select 
		id,
		name,
		description,
		price,
		created_by
		from Products;
	`

	qryGetProductsByID = `
		select 
		id,
		name,
		description,
		price,
		created_by
		from Products
		where id = $;
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInserProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}
	err := r.db.SelectContext(ctx, &pp, qryGetAllProducts) //obtiene multiples records
	if err != nil {
		return nil, err
	}
	return pp, err
}
func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, qryGetProductsByID, id) //obtiene 1 record
	if err != nil {
		return nil, err
	}
	return p, nil
}
