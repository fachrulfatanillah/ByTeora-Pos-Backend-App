package route

import (
	"ByTeora-Pos-Backend-App/controller"
	"ByTeora-Pos-Backend-App/middleware"

	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {

	// Default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Byteora API",
		})
	})

	api := r.Group("/api/v1")

	// ---------- USER ----------
	api.POST("/users", controller.CreateUser)

	// ---------- AUTH ----------
	api.POST("/auth/login", controller.AuthLogin)

	// ---------- STORE ----------
	store := api.Group("/store")
	store.Use(middleware.AuthMiddleware())
	{
		store.POST("/", controller.CreateStore)
		store.POST("/list", controller.GetStoresByUser)
		store.PUT("/:store_uuid", controller.UpdateStore)
		store.DELETE("/:store_uuid", controller.DeleteStore)
	}

	// ---------- CATEGORY ----------
	category := api.Group("/stores/:store_uuid/categories")
	category.Use(middleware.AuthMiddleware())
	{
		category.POST("/", controller.CreateCategory)
		category.GET("/", controller.GetCategoriesByStore)
		category.PUT("/:category_uuid", controller.UpdateCategory)
		category.DELETE("/:category_uuid", controller.DeleteCategory)
	}

	// ---------- PRODUCT ----------
	product := api.Group("/stores/:store_uuid/products")
	product.Use(middleware.AuthMiddleware())
	{
		product.POST("/", controller.CreateProduct)
		product.GET("/", controller.GetAllProducts)
		product.PUT("/:product_uuid", controller.UpdateProductHandler)
		product.DELETE("/:product_uuid", controller.DeleteProductHandler)
	}
	

	// ---------- PRODUCT STOCKS ----------
	product_stocks := api.Group("/stores/:store_uuid/products")
	product_stocks.Use(middleware.AuthMiddleware())
	{
		product_stocks.POST("/:product_uuid/stocks/", controller.CreateProductStockHandler)
		product_stocks.GET("/stocks/log", controller.GetAllProductStockLogsHandler)
		product_stocks.GET("/:product_uuid/stocks/log", controller.GetProductStockLogsByProductHandler)
		product_stocks.GET("/stocks", controller.GetAllProductStocksHandler)
		product_stocks.GET("/:product_uuid/stocks", controller.GetProductCurrentStockHandler)
	}
}