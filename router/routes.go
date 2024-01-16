package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leocardhio/ecom-catalogue/controller"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/repository"
	"github.com/leocardhio/ecom-catalogue/service"
)

type Controllers struct {
	ProductController controller.IProductController
}

func NewRouter(dbs db.Database) *gin.Engine {
	productRepository := repository.NewProductRepository(dbs)

	productService := service.NewProductService(productRepository)

	controllers := Controllers{
		ProductController: controller.NewProductController(productService),
	}

	router := gin.New()
	setMiddleware(router)
	setRouter(router, controllers)

	return router
}

func setMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func setRouter(r *gin.Engine, controllers Controllers) {
	ver := r.Group("/api/v1")
	ver.POST("/products", controllers.ProductController.CreateProduct)
	ver.GET("/products/:id", controllers.ProductController.GetProduct)
	ver.GET("/products", controllers.ProductController.GetProducts)
	ver.PUT("/products/:id", controllers.ProductController.UpdateProduct)
	ver.PATCH("/products/:id/status", controllers.ProductController.UpdateProductStatus)
	ver.DELETE("/products/:id", controllers.ProductController.DeleteProduct)
}
