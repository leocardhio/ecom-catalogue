package service

import (
	"context"
	"errors"

	"github.com/leocardhio/ecom-catalogue/datastruct"
	"github.com/leocardhio/ecom-catalogue/db/repository"
	"github.com/leocardhio/ecom-catalogue/model"
	"github.com/leocardhio/ecom-catalogue/util"
)

const (
	ErrInvalidTags = "invalid tags"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.CreateProductResponse, error)
	GetProduct(ctx context.Context, req model.GetProductRequest) (datastruct.Product, error)
	GetProducts(ctx context.Context, req model.GetProductsRequest) (model.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req model.UpdateProductRequest) (model.UpdateProductResponse, error)
	UpdateProductStatus(ctx context.Context, req model.UpdateProductStatusRequest) (model.UpdateProductStatusResponse, error)
	DeleteProduct(ctx context.Context, req model.DeleteProductRequest) (model.DeleteProductResponse, error)
}

type productService struct {
	productRepository repository.IProductRepository
	tagRepository     repository.ITagsRepository
}

func (service *productService) getUpdateTagsCommand(productId string, updatedTags []datastruct.Tag) (datastruct.UpdateTagMap, error) {
	tagsMap := datastruct.UpdateTagMap{}

	queryRes, err := service.tagRepository.GetTagsByProductId(context.Background(), repository.GetTagsByProductIdParams{
		ProductId: productId,
	})
	if err != nil {
		return nil, err
	}

	for _, res := range queryRes {
		tagsMap[res.TagId] = datastruct.REMOVE_TAG
	}

	for _, tag := range updatedTags {
		if _, ok := tagsMap[tag.Id]; !ok {
			tagsMap[tag.Id] = datastruct.ADD_TAG
		} else {
			delete(tagsMap, tag.Id)
		}
	}

	return tagsMap, nil
}

func (service *productService) validateTags(productTags []datastruct.Tag) error {
	tags, err := service.tagRepository.GetTags(context.Background())
	if err != nil {
		return err
	}

	tagTable := map[string]struct{}{}
	for _, tag := range tags {
		tagTable[tag.Name] = struct{}{}
	}

	for _, tag := range productTags {
		if _, ok := tagTable[tag.Name]; !ok {
			return errors.New(ErrInvalidTags)
		}
	}
	return nil
}

func NewProductService(productRepository repository.IProductRepository, tagRepository repository.ITagsRepository) IProductService {
	return &productService{
		productRepository: productRepository,
		tagRepository:     tagRepository,
	}
}

func (service *productService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.CreateProductResponse, error) {
	var res model.CreateProductResponse
	var err error

	ulid := util.GetUlid()

	err = service.validateTags(req.Tags)
	if err != nil {
		return res, err
	}

	res.Count, err = service.productRepository.CreateProduct(ctx, repository.CreateProductParams{
		Id:          ulid,
		Name:        req.Name,
		Price:       req.Price,
		Tags:        req.Tags,
		Description: req.Description,
		Condition:   req.Condition,
	})
	if err != nil {
		return res, err
	}

	res.Id = ulid
	return res, nil
}

func (service *productService) GetProduct(ctx context.Context, req model.GetProductRequest) (datastruct.Product, error) {
	var res datastruct.Product

	// TODO: To be implemented
	return res, nil
}

func (service *productService) GetProducts(ctx context.Context, req model.GetProductsRequest) (model.GetProductsResponse, error) {
	var res model.GetProductsResponse

	// TODO: To be implemented
	return res, nil
}

func (service *productService) UpdateProduct(ctx context.Context, req model.UpdateProductRequest) (model.UpdateProductResponse, error) {
	var res model.UpdateProductResponse
	var err error

	if err = service.validateTags(req.Tags); err != nil {
		return res, err
	}

	updateTagCommands, err := service.getUpdateTagsCommand(req.Id, req.Tags)
	if err != nil {
		return res, err
	}

	res.Count, err = service.productRepository.UpdateProduct(ctx, repository.UpdateProductParams{
		Id:          req.Id,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Condition:   req.Condition,
		Commands:    updateTagCommands,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *productService) UpdateProductStatus(ctx context.Context, req model.UpdateProductStatusRequest) (model.UpdateProductStatusResponse, error) {
	var res model.UpdateProductStatusResponse
	var err error

	res.Count, err = service.productRepository.UpdateProductStatus(ctx, repository.UpdateProductStatusParams{
		Id:     req.Id,
		IsSold: req.IsSold,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *productService) DeleteProduct(ctx context.Context, req model.DeleteProductRequest) (model.DeleteProductResponse, error) {
	var res model.DeleteProductResponse
	var err error

	res.Count, err = service.productRepository.DeleteProduct(ctx, repository.DeleteProductParams{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
