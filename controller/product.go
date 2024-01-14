package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leocardhio/ecom-catalogue/db/model"
	"github.com/leocardhio/ecom-catalogue/service"
	"github.com/leocardhio/ecom-catalogue/util"
)

type IProductController interface {
	CreateProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productController struct {
	productService service.IProductService
}

func NewProductController(service service.IProductService) IProductController {
	return &productController{
		productService: service,
	}
}

func (controller *productController) CreateProduct(c *gin.Context) {
	var product model.CreateProductRequest

	if err := c.ShouldBindJSON(&product); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.productService.CreateProduct(c, product)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseCreated(c, res)
}

func (controller *productController) GetProduct(c *gin.Context) {
	// TODO: Implement me
}

func (controller *productController) GetProducts(c *gin.Context) {
	// TODO: Implement me
}

func (controller *productController) UpdateProduct(c *gin.Context) {
	// TODO: Implement me
}

func (controller *productController) DeleteProduct(c *gin.Context) {
	// TODO: Implement me
}	