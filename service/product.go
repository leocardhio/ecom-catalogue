package service

import (
	"context"

	"github.com/leocardhio/ecom-catalogue/class"
	"github.com/leocardhio/ecom-catalogue/db/model"
	"github.com/leocardhio/ecom-catalogue/db/repository"
	"github.com/oklog/ulid"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (model.CreateProductResponse, error)
	GetProduct(ctx context.Context, req model.GetProductRequest) (model.ProductModel, error)
	GetProducts(ctx context.Context, req model.GetProductsRequest) (model.GetProductsResponse, error)
	UpdateProduct(ctx context.Context, req model.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, req model.DeleteProductRequest) error
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
	ulid := ulid.MustNew(ulid.Now(), nil)
	
	product, err := service.productRepository.CreateProduct(ctx, class.Product{
		Ulid:        ulid.String(),
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Condition:   req.Condition,
	})
	if err!= nil {
		return res, err
	}

	res = model.CreateProductResponse{
		Ulid:        product.Ulid,
	}

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

func (service *productService) UpdateProduct(ctx context.Context, req model.UpdateProductRequest) error {	
	// TODO: To be implemented
	return nil
}

func (service *productService) DeleteProduct(ctx context.Context, req model.DeleteProductRequest) error {
	// TODO: To be implemented
	return nil
}

