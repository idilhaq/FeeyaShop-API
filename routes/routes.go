package routes

import (
	"feeyashop/controller"
	"feeyashop/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Category
	r.GET("/category", controller.GetAllCategory)
	r.GET("/category/:id", controller.GetProductsByCategoryId)
	categoryMiddlewareRoute := r.Group("/category")
	categoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	categoryMiddlewareRoute.POST("/", controller.CreateCategory)
	categoryMiddlewareRoute.PATCH("/:id", controller.UpdateCategory)
	categoryMiddlewareRoute.DELETE("/:id", controller.DeleteCategory)

	// Product
	r.GET("/product", controller.GetAllProduct)
	productMiddlewareRoute := r.Group("/product")
	productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	productMiddlewareRoute.POST("/", controller.CreateProduct)
	productMiddlewareRoute.PATCH("/:id", controller.UpdateProduct)
	productMiddlewareRoute.DELETE("/:id", controller.DeleteProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
