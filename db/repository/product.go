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
	CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error)
	GetProduct(ctx context.Context, arg GetProductParams) (class.Product, error)
	GetProducts(ctx context.Context, arg GetProductsParams) ([]class.Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error)
	DeleteProduct(ctx context.Context, arg DeleteProductParams) (int64, error)
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

func (repo *productRepository) CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error) {
	var count int64

	result, err := repo.writeDB.ExecContext(ctx, query.CreateProduct, arg.Id, arg.Name, arg.Price, arg.Description, arg.Condition)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
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

func (repo *productRepository) UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error) {
	var count int64

	result, err := repo.writeDB.ExecContext(ctx, query.UpdateProduct, arg.Name, arg.Price, arg.Description, arg.Condition, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}

type DeleteProductParams struct {
	Id string
}

func (repo *productRepository) DeleteProduct(ctx context.Context, arg DeleteProductParams) (int64, error) {
	var count int64

	result, err := repo.writeDB.ExecContext(ctx, query.DeleteProduct, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}
