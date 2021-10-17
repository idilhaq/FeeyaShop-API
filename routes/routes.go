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
	categoryMiddlewareRoute := r.Group("/category")
	categoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	categoryMiddlewareRoute.POST("/", controller.CreateCategory)
	categoryMiddlewareRoute.PATCH("/:id", controller.UpdateCategory)
	categoryMiddlewareRoute.DELETE("/:id", controller.DeleteCategory)

	// Product
	// r.GET("/product", controllers.GetAllProduct)
	// r.POST("/product/create", controllers.PostProduct)
	// r.PUT("/product/:id/update", controllers.UpdateProduct)
	// r.DELETE("/product/:id/delete", controllers.DeleteProduct)

	// productMiddlewareRoute := r.Group("/product")
	// productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	// productMiddlewareRoute.POST("/", controllers.CreateProduct)
	// productMiddlewareRoute.PATCH("/:id", controllers.UpdateProduct)
	// productMiddlewareRoute.DELETE("/:id", controllers.DeleteProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
