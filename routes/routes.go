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
	userMiddlewareRoute := r.Group("/user")
	userMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	userMiddlewareRoute.PATCH("/:id", controller.ChangePassword)

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

	// Role
	r.GET("/role", controller.GetAllRole)
	roleMiddlewareRoute := r.Group("/role")
	roleMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	roleMiddlewareRoute.POST("/", controller.CreateRole)
	roleMiddlewareRoute.PATCH("/:id", controller.UpdateRole)
	roleMiddlewareRoute.DELETE("/:id", controller.DeleteRole)

	// Like
	r.GET("/like", controller.GetAllLike)
	likeMiddlewareRoute := r.Group("/like")
	likeMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	likeMiddlewareRoute.POST("/", controller.CreateLike)
	likeMiddlewareRoute.DELETE("/:id", controller.DeleteLike)

	// Cart
	r.GET("/cart", controller.GetAllCart)
	cartMiddlewareRoute := r.Group("/cart")
	cartMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	cartMiddlewareRoute.POST("/", controller.CreateCart)
	cartMiddlewareRoute.PATCH("/:id", controller.UpdateCart)
	cartMiddlewareRoute.DELETE("/:id", controller.DeleteCart)

	// Purchase
	r.GET("/purchase", controller.GetAllPurchase)
	purchaseMiddlewareRoute := r.Group("/purchase")
	purchaseMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	purchaseMiddlewareRoute.POST("/", controller.CreatePurchase)
	purchaseMiddlewareRoute.PATCH("/:id", controller.UpdatePurchase)
	purchaseMiddlewareRoute.DELETE("/:id", controller.DeletePurchase)

	// Rating
	r.GET("/rating", controller.GetAllRating)
	ratingMiddlewareRoute := r.Group("/rating")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controller.CreateRating)
	ratingMiddlewareRoute.PATCH("/:id", controller.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controller.DeleteRating)

	// Comment
	r.GET("/comment", controller.GetAllComment)
	commentMiddlewareRoute := r.Group("/comment")
	commentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	commentMiddlewareRoute.POST("/", controller.CreateComment)
	commentMiddlewareRoute.PATCH("/:id", controller.UpdateComment)
	commentMiddlewareRoute.DELETE("/:id", controller.DeleteComment)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
