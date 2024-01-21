package repository

import (
	"context"
	"database/sql"

	"github.com/leocardhio/ecom-catalogue/datastruct"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/query"
)

type ITagsRepository interface {
	CreateTag(ctx context.Context, arg CreateTagParams) (CreateTagResult, error)
	GetTags(ctx context.Context) ([]datastruct.Tag, error)
	GetTagsByProductId(ctx context.Context, arg GetTagsByProductIdParams) ([]GetTagsByProductIdResult, error)
	GetTag(ctx context.Context, arg GetTagParams) (*datastruct.Tag, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (int64, error)
	DeleteTag(ctx context.Context, arg DeleteTagParams) (int64, error)
}

type tagsRepository struct {
	db db.Database
}

func NewTagsRepository(dbs db.Database) ITagsRepository {
	return &tagsRepository{
		db: dbs,
	}
}

type CreateTagParams struct {
	Name string
}

type CreateTagResult struct {
	Id    int64
	Count int64
}

func (repo *tagsRepository) CreateTag(ctx context.Context, arg CreateTagParams) (CreateTagResult, error) {
	var res CreateTagResult

	row := repo.db.GetPrimary().QueryRowContext(ctx, query.CreateTag, arg.Name)

	if err := row.Scan(&res.Id); err != nil {
		return res, err
	}
	res.Count += 1

	return res, nil
}

func (repo *tagsRepository) GetTags(ctx context.Context) ([]datastruct.Tag, error) {
	var res []datastruct.Tag

	rows, err := repo.db.GetPrimary().QueryContext(ctx, query.GetTags)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var tag datastruct.Tag
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return res, err
		}
		res = append(res, tag)
	}

	return res, nil
}

type GetTagsByProductIdParams struct {
	ProductId string
}

type GetTagsByProductIdResult struct {
	TagId string
}

func (repo *tagsRepository) GetTagsByProductId(ctx context.Context, arg GetTagsByProductIdParams) ([]GetTagsByProductIdResult, error) {
	var res []GetTagsByProductIdResult

	rows, err := repo.db.GetPrimary().QueryContext(ctx, query.GetTagsByProductId, arg.ProductId)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var result GetTagsByProductIdResult
		if err := rows.Scan(&result.TagId); err != nil {
			return res, err
		}
		res = append(res, result)
	}

	return res, nil
}

type GetTagParams struct {
	Id string
}

func (repo *tagsRepository) GetTag(ctx context.Context, arg GetTagParams) (*datastruct.Tag, error) {
	var res datastruct.Tag

	row := repo.db.GetPrimary().QueryRowContext(ctx, query.GetTag, arg.Id)
	if err := row.Scan(&res.Id, &res.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return &res, err
		}
	}

	return &res, nil
}

type UpdateTagParams struct {
	Id   string
	Name string
}

func (repo *tagsRepository) UpdateTag(ctx context.Context, arg UpdateTagParams) (int64, error) {
	var count int64

	result, err := repo.db.GetPrimary().ExecContext(ctx, query.UpdateTag, arg.Name, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}

type DeleteTagParams struct {
	Id string
}

func (repo *tagsRepository) DeleteTag(ctx context.Context, arg DeleteTagParams) (int64, error) {
	var count int64

	result, err := repo.db.GetPrimary().ExecContext(ctx, query.DeleteTag, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}
