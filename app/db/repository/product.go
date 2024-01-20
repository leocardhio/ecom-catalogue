package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/leocardhio/ecom-catalogue/datastruct"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/query"
	"github.com/leocardhio/ecom-catalogue/util"
)

const (
	ErrInvalidCommand = "invalid command"
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
	db db.Database
}

func NewProductRepository(dbs db.Database) IProductRepository {
	return &productRepository{
		db: dbs,
	}
}

type CreateProductParams struct {
	Id          string
	Name        string
	Price       uint
	Tags        []datastruct.Tag
	Description string
	Condition   datastruct.ProductCondition
}

func (repo *productRepository) CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error) {
	var countTotal int64

	util.ExecTx(ctx, repo.db.GetPrimary(), func(tx *sql.Tx) error {
		var err error
		result, err := tx.ExecContext(ctx, query.CreateProduct, arg.Id, arg.Name, arg.Price, arg.Description, arg.Condition)
		if err != nil {
			return err
		}

		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		countTotal += count

		for _, tag := range arg.Tags {
			_, err = tx.ExecContext(ctx, query.CreateProductTags, arg.Id, tag.Id)
			if err != nil {
				return err
			}
			count, err = result.RowsAffected()
			if err != nil {
				return err
			}
			countTotal += count
		}

		return nil
	})

	return countTotal, nil
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
	Commands    datastruct.UpdateTagMap
}

func (repo *productRepository) UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error) {
	// TODO: Apply Tx for Tags to this method
	var countTotal int64

	util.ExecTx(ctx, repo.db.GetPrimary(), func(tx *sql.Tx) error {
		for tid, command := range arg.Commands {
			switch datastruct.UpdateTagCommand(command) {
			case datastruct.ADD_TAG:
				result, err := tx.ExecContext(ctx, query.CreateProductTags, arg.Id, tid)
				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			case datastruct.REMOVE_TAG:
				result, err := tx.ExecContext(ctx, query.DeleteProductTags, arg.Id, tid)
				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			default:
				return errors.New(ErrInvalidCommand)
			}
		}

		return nil
	})

	return countTotal, nil
}

type UpdateProductStatusParams struct {
	Id     string
	IsSold bool
}

func (repo *productRepository) UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) (int64, error) {
	var count int64

	result, err := repo.db.GetPrimary().ExecContext(ctx, query.UpdateProductStatus, arg.IsSold, arg.Id)
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

	result, err := repo.db.GetPrimary().ExecContext(ctx, query.DeleteProduct, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}
