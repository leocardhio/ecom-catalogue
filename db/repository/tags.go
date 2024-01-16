package repository

import (
	"context"
	"database/sql"

	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/query"
)

type ITagsRepository interface {
	CreateTag(ctx context.Context, arg CreateTagParams) (CreateTagResult, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (int64, error)
	DeleteTag(ctx context.Context, arg DeleteTagParams) (int64, error)
}

type tagsRepository struct {
	writeDB *sql.DB
}

func NewTagsRepository(dbs db.Database) ITagsRepository {
	return &tagsRepository{
		writeDB: dbs.GetSQL(),
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

	result, err := repo.writeDB.ExecContext(ctx, query.CreateTag, arg.Name)
	if err != nil {
		return res, err
	}

	if res.Count, err = result.RowsAffected(); err != nil {
		return res, err
	}

	if res.Id, err = result.LastInsertId(); err != nil {
		return res, err
	}

	return res, nil
}

type UpdateTagParams struct {
	Id   string
	Name string
}

func (repo *tagsRepository) UpdateTag(ctx context.Context, arg UpdateTagParams) (int64, error) {
	var count int64

	result, err := repo.writeDB.ExecContext(ctx, query.UpdateTag, arg.Name, arg.Id)
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

	result, err := repo.writeDB.ExecContext(ctx, query.DeleteTag, arg.Id)
	if err != nil {
		return count, err
	}

	if count, err = result.RowsAffected(); err != nil {
		return count, err
	}

	return count, nil
}
