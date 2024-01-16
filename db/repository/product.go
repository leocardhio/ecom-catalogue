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
	CreateProduct(ctx context.Context, arg CreateProductParams) (class.Product, error)
	GetProduct(ctx context.Context, arg GetProductParams) (class.Product, error)
	GetProducts(ctx context.Context, arg GetProductsParams) ([]class.Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (class.Product, error)
	DeleteProduct(ctx context.Context, arg DeleteProductParams) error
}

type productRepository struct {
	readDB  *elasticsearch.Client
	writeDB *sql.DB
}

func NewProductRepository(dbs db.Database) IProductRepository {
	return &productRepository{
		readDB:  dbs.GetES(),
		writeDB: dbs.GetSQL(),
	}
}

type CreateProductParams struct {
	Id          string
	Name        string
	Price       uint
	Description string
	Condition   class.ProductCondition
}

func (repo *productRepository) CreateProduct(ctx context.Context, arg CreateProductParams) (class.Product, error) {
	// TODO: Remove RETURNING clause
	var createdProduct class.Product

	row, err := repo.writeDB.QueryContext(ctx, query.CreateProduct, arg.Id, arg.Name, arg.Price, arg.Description, arg.Condition)
	if err != nil {
		return createdProduct, err
	}
	defer row.Close()

	row.Next()
	if err = row.Err(); err != nil {
		return createdProduct, err
	}

	if err = row.Scan(
		&createdProduct.Id,
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

type GetProductParams struct {
	Id string
}

func (repo *productRepository) GetProduct(ctx context.Context, arg GetProductParams) (class.Product, error) {
	// TODO: To be implemented
	return class.Product{}, nil
}

type GetProductsParams struct {
	Limit  uint8
	Offset uint
}

func (repo *productRepository) GetProducts(ctx context.Context, arg GetProductsParams) ([]class.Product, error) {
	// TODO: To be implemented
	return []class.Product{}, nil
}

type UpdateProductParams struct {
	Id          string
	Name        string
	Price       uint
	Description string
	Condition   class.ProductCondition
}

func (repo *productRepository) UpdateProduct(ctx context.Context, arg UpdateProductParams) (class.Product, error) {
	// TODO: Remove RETURNING clause
	var updatedProduct class.Product

	row, err := repo.writeDB.QueryContext(ctx, query.UpdateProduct, arg.Name, arg.Price, arg.Description, arg.Condition, arg.Id)
	if err != nil {
		return updatedProduct, err
	}
	defer row.Close()

	row.Next()
	if err := row.Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Description,
		&updatedProduct.Condition,
		&updatedProduct.UpdatedAt,
		&updatedProduct.DeletedAt,
	); err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

type DeleteProductParams struct {
	Id string
}

func (repo *productRepository) DeleteProduct(ctx context.Context, arg DeleteProductParams) error {
	// TODO: To be implemented
	return nil
}
