package service

import (
	"context"

	"github.com/leocardhio/ecom-catalogue/db/model"
	"github.com/leocardhio/ecom-catalogue/db/repository"
	"github.com/leocardhio/ecom-catalogue/util"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.CreateProductResponse, error)
	GetProduct(ctx context.Context, req model.GetProductRequest) (model.ProductModel, error)
	GetProducts(ctx context.Context, req model.GetProductsRequest) (model.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req model.UpdateProductRequest) (model.UpdateProductResponse, error)
	UpdateProductStatus(ctx context.Context, req model.UpdateProductStatusRequest) (model.UpdateProductStatusResponse, error)
	DeleteProduct(ctx context.Context, req model.DeleteProductRequest) (model.DeleteProductResponse, error)
}

type productService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (service *productService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.CreateProductResponse, error) {
	var res model.CreateProductResponse
	var err error

	ulid := util.GetUlid()
	res.Count, err = service.productRepository.CreateProduct(ctx, repository.CreateProductParams{
		Id:          ulid,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Condition:   req.Condition,
	})
	if err != nil {
		return res, err
	}

	res.Id = ulid
	return res, nil
}

func (service *productService) GetProduct(ctx context.Context, req model.GetProductRequest) (model.ProductModel, error) {
	var res model.ProductModel

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

	res.Count, err = service.productRepository.UpdateProduct(ctx, repository.UpdateProductParams{
		Id:          req.Id,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Condition:   req.Condition,
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
