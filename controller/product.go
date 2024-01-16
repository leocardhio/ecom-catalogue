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
	UpdateProductStatus(c *gin.Context)
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
	var req model.CreateProductRequest
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.productService.CreateProduct(c, req)
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
	var req model.UpdateProductRequest
	var err error

	err = c.ShouldBindUri(&req.UpdateProductRequestURI)
	if err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	err = c.ShouldBindJSON(&req.UpdateProductRequestBody)
	if err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.productService.UpdateProduct(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseUpdated(c, res)
}

func (controller *productController) UpdateProductStatus(c *gin.Context) {
	var req model.UpdateProductStatusRequest
	var err error

	if err = c.ShouldBindUri(&req.UpdateProductStatusRequestURI); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	if err = c.ShouldBindJSON(&req.UpdateProductStatusRequestBody); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.productService.UpdateProductStatus(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseUpdated(c, res)
}

func (controller *productController) DeleteProduct(c *gin.Context) {
	var req model.DeleteProductRequest
	var err error

	if err = c.ShouldBindUri(&req); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.productService.DeleteProduct(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseDeleted(c, res)
}
