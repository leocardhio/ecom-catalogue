package repository

import (
	"context"
	"database/sql"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/leocardhio/ecom-catalogue/class"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/query"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product class.Product) (class.Product, error)
	GetProduct(ctx context.Context, ulid string) (class.Product, error)
	GetProducts(ctx context.Context) ([]class.Product, error)
	UpdateProduct(ctx context.Context, product class.Product) (class.Product, error)
	DeleteProduct(ctx context.Context, ulid string) error
}

type productRepository struct {
	readDB *elasticsearch.Client
	writeDB *sql.DB
}

func NewProductRepository(dbs db.Database) IProductRepository {
	return &productRepository{
		readDB: dbs.GetES(),
		writeDB: dbs.GetSQL(),
	}
}

func (repo *productRepository) CreateProduct(ctx context.Context, product class.Product) (class.Product, error) {
	// TODO: Remove RETURNING clause
	var createdProduct class.Product

	row, err := repo.writeDB.QueryContext(ctx, query.CreateProduct, product.Ulid, product.Name, product.Price, product.Description, product.Condition)
	if err != nil {
		return createdProduct, err
	}
	defer row.Close()
	
	row.Next() 
	if err = row.Err(); err != nil { return createdProduct, err }

	if err = row.Scan(
		&createdProduct.Ulid, 
		&createdProduct.Name, 
		&createdProduct.Price, 
		&createdProduct.Description, 
		&createdProduct.Condition, 
		&createdProduct.UpdatedAt, 
		&createdProduct.DeletedAt,
	); err != nil {
		return createdProduct, err
	}

	return createdProduct, nil
}

func (repo *productRepository) GetProduct(ctx context.Context, ulid string) (class.Product, error) {
	// TODO: To be implemented
	return class.Product{}, nil
}

func (repo *productRepository) GetProducts(ctx context.Context) ([]class.Product, error) {
	// TODO: To be implemented
	return []class.Product{}, nil
}

func (repo *productRepository) UpdateProduct(ctx context.Context, product class.Product) (class.Product, error) {
	// TODO: To be implemented
	return class.Product{}, nil
}

func (repo *productRepository) DeleteProduct(ctx context.Context, ulid string) error {
	// TODO: To be implemented
	return nil
}