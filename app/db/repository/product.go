package repository

import (
	"context"
	"database/sql"

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
	GetProductImageUrlsByProductId(ctx context.Context, arg GetImageUrlsByProductIdParams) (GetImageUrlsByProductIdResult, error)
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
	ImageUrls   []string
	Description string
	Condition   datastruct.ProductCondition
}

func (repo *productRepository) CreateProduct(ctx context.Context, arg CreateProductParams) (int64, error) {
	var countTotal int64

	err := util.ExecTx(ctx, repo.db.GetPrimary(), func(tx *sql.Tx) error {
		var err error
		var result sql.Result

		//? Image handler
		for _, imageUrl := range arg.ImageUrls {
			result, err = tx.ExecContext(ctx, query.CreateProductImage, arg.Id, imageUrl)
			if err != nil {
				return err
			}

			count, err := result.RowsAffected()
			if err != nil {
				return err
			}
			countTotal += count
		}

		//? Tags handler
		result, err = tx.ExecContext(ctx, query.CreateProduct, arg.Id, arg.Name, arg.Price, arg.Description, arg.Condition)
		if err != nil {
			return err
		}

		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		countTotal += count

		for _, tag := range arg.Tags {
			result, err = tx.ExecContext(ctx, query.CreateProductTags, arg.Id, tag.Id)
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

	if err != nil {
		return countTotal, err
	}

	return countTotal, nil
}

type GetImageUrlsByProductIdParams struct {
	Id string
}

type GetImageUrlsByProductIdResult struct {
	ImageUrls []string
}

func (repo *productRepository) GetProductImageUrlsByProductId(ctx context.Context, arg GetImageUrlsByProductIdParams) (GetImageUrlsByProductIdResult, error) {
	var res GetImageUrlsByProductIdResult

	rows, err := repo.db.GetPrimary().QueryContext(ctx, query.GetProductImageUrlsByProductId, arg.Id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var imageUrl string
		if err := rows.Scan(&imageUrl); err != nil {
			return res, err
		}

		res.ImageUrls = append(res.ImageUrls, imageUrl)
	}

	return res, nil
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
	Id               string
	Name             string
	Price            uint
	Description      string
	Condition        datastruct.ProductCondition
	TagCommands      datastruct.UpdateTagMap
	ImageUrlCommands datastruct.UpdateImageUrlMap
}

func (repo *productRepository) UpdateProduct(ctx context.Context, arg UpdateProductParams) (int64, error) {
	var countTotal int64

	err := util.ExecTx(ctx, repo.db.GetPrimary(), func(tx *sql.Tx) error {
		var err error
		var result sql.Result

		for url, command := range arg.ImageUrlCommands {
			if datastruct.UpdateImageUrlCommand(command) == datastruct.ADD_IMAGE_URL {
				result, err = tx.ExecContext(ctx, query.CreateProductImage, arg.Id, url)
				if err != nil {
					return err
				}

				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			} else {
				result, err = tx.ExecContext(ctx, query.DeleteProductImage, arg.Id, url)
				if err != nil {
					return err
				}

				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			}
		}

		result, err = tx.ExecContext(ctx, query.UpdateProduct, arg.Name, arg.Price, arg.Description, arg.Condition, arg.Id)
		if err != nil {
			return err
		}

		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		countTotal += count

		for tid, command := range arg.TagCommands {
			if datastruct.UpdateTagCommand(command) == datastruct.ADD_TAG {
				result, err := tx.ExecContext(ctx, query.CreateProductTags, arg.Id, tid)
				if err != nil {
					return err
				}

				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			} else {
				result, err := tx.ExecContext(ctx, query.DeleteProductTags, arg.Id, tid)
				if err != nil {
					return err
				}

				count, err := result.RowsAffected()
				if err != nil {
					return err
				}
				countTotal += count
			}
		}
		return nil
	})

	if err != nil {
		return countTotal, err
	}

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
