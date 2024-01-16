package service

import (
	"context"
	"fmt"

	"github.com/leocardhio/ecom-catalogue/db/model"
	"github.com/leocardhio/ecom-catalogue/db/repository"
)

type ITagsService interface {
	CreateTag(ctx context.Context, req model.CreateTagRequest) (model.CreateTagResponse, error)
	UpdateTag(ctx context.Context, req model.UpdateTagRequest) (model.UpdateTagResponse, error)
	DeleteTag(ctx context.Context, req model.DeleteTagRequest) (model.DeleteTagResponse, error)
}

type tagsService struct {
	tagRepository repository.ITagsRepository
}

func NewTagsService(tagRepository repository.ITagsRepository) ITagsService {
	return &tagsService{
		tagRepository: tagRepository,
	}
}

func (service *tagsService) CreateTag(ctx context.Context, req model.CreateTagRequest) (model.CreateTagResponse, error) {
	var res model.CreateTagResponse
	var err error

	result, err := service.tagRepository.CreateTag(ctx, repository.CreateTagParams{
		Name: req.Name,
	})
	if err != nil {
		return res, err
	}

	res.Count = result.Count
	res.Id = fmt.Sprint(result.Id)
	return res, nil
}

func (service *tagsService) UpdateTag(ctx context.Context, req model.UpdateTagRequest) (model.UpdateTagResponse, error) {
	var res model.UpdateTagResponse
	var err error

	result, err := service.tagRepository.UpdateTag(ctx, repository.UpdateTagParams{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return res, err
	}

	res.Count = result
	return res, nil
}

func (service *tagsService) DeleteTag(ctx context.Context, req model.DeleteTagRequest) (model.DeleteTagResponse, error) {
	var res model.DeleteTagResponse
	var err error

	result, err := service.tagRepository.DeleteTag(ctx, repository.DeleteTagParams{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res.Count = result
	return res, nil
}
