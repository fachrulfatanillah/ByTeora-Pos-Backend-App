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
		}
	}