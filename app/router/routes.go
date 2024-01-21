package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leocardhio/ecom-catalogue/controller"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/db/repository"
	"github.com/leocardhio/ecom-catalogue/middleware"
	"github.com/leocardhio/ecom-catalogue/service"
)

type Controllers struct {
	ProductController controller.IProductController
	TagsController    controller.ITagsController
}

func NewRouter(dbs db.Database) *gin.Engine {
	productRepository := repository.NewProductRepository(dbs)
	tagsRepository := repository.NewTagsRepository(dbs)

	productService := service.NewProductService(productRepository, tagsRepository)
	tagsService := service.NewTagsService(tagsRepository)

	controllers := Controllers{
		ProductController: controller.NewProductController(productService),
		TagsController:    controller.NewTagsController(tagsService),
	}

	router := gin.New()
	setMiddleware(router)
	setRouter(router, controllers)

	return router
}

func setMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORS())
}

func setRouter(r *gin.Engine, controllers Controllers) {
	ver := r.Group("/api/v1")

	product := ver.Group("/products")
	product.POST("", controllers.ProductController.CreateProduct)
	product.GET("/:id", controllers.ProductController.GetProduct)
	product.GET("", controllers.ProductController.GetProducts)
	product.PUT("/:id", controllers.ProductController.UpdateProduct)
	product.PATCH("/:id/status", controllers.ProductController.UpdateProductStatus)
	product.DELETE("/:id", controllers.ProductController.DeleteProduct)

	tag := ver.Group("/tags")
	tag.POST("", controllers.TagsController.CreateTag)
	tag.DELETE("/:id", controllers.TagsController.DeleteTag)
	tag.PUT("/:id", controllers.TagsController.UpdateTag)
}
