package repository

import (
	"context"
	"database/sql"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/leocardhio/ecom-catalogue/datastruct"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/query"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error)
	GetProduct(ctx context.Context, arg GetProductParams) (datastruct.Product, error)
	GetProducts(ctx context.Context, arg GetProductsParams) ([]datastruct.Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error)
	UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) (int64, error)
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
	Condition   datastruct.ProductCondition
}

func (repo *productRepository) validateTags(tags []datastruct.Tag) bool {
	// TODO: To be implemented
	return true
}

func (repo *productRepository) CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error) {
	// TODO: Apply Tx for Tags to this method
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

func (repo *productRepository) GetProduct(ctx context.Context, arg GetProductParams) (datastruct.Product, error) {
	// TODO: To be implemented
	return datastruct.Product{}, nil
}

type GetProductsParams struct {
	Limit  uint8
	Offset uint
}

func (repo *productRepository) GetProducts(ctx context.Context, arg GetProductsParams) ([]datastruct.Product, error) {
	// TODO: To be implemented
	return []datastruct.Product{}, nil
}

type UpdateProductParams struct {
	Id          string
	Name        string
	Price       uint
	Description string
	Condition   datastruct.ProductCondition
}

func (repo *productRepository) UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error) {
	// TODO: Apply Tx for Tags to this method
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

type UpdateProductStatusParams struct {
	Id     string
	IsSold bool
}

func (repo *productRepository) UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) (int64, error) {
	var count int64

	result, err := repo.writeDB.ExecContext(ctx, query.UpdateProductStatus, arg.IsSold, arg.Id)
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
