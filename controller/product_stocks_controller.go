package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ByTeora-Pos-Backend-App/service"
	"ByTeora-Pos-Backend-App/api/request"
)

func CreateProductStockHandler(c *gin.Context) {
	storeUUID := c.Param("store_uuid")
	productUUID := c.Param("product_uuid")

	if storeUUID == "" || productUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "store_uuid and product_uuid are required",
		})
		return
	}

	userUUID, exists := c.Get("user_uuid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	owned, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
	if err != nil || !owned {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": "You are not allowed to access this store",
		})
		return
	}

	storeID, err := service.GetStoreIDByUUID(storeUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Store not found",
		})
		return
	}

	productID, err := service.GetProductIDByUUID(productUUID, storeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Product not found",
		})
		return
	}

	var req request.CreateProductStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON",
		})
		return
	}

	stock, err := service.CreateProductStock(storeID, productID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create stock",
		})
		return
	}

	stock.ProductUUID = productUUID
	stock.StoreUUID = storeUUID

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Stock created successfully",
		"data":    stock,
	})
}

func GetAllProductStockLogsHandler(c *gin.Context) {
    storeUUID := c.Param("store_uuid")

    userUUID, exists := c.Get("user_uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Unauthorized",
        })
        return
    }

    owned, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    if !owned {
        c.JSON(http.StatusForbidden, gin.H{
            "status":  "error",
            "message": "You are not allowed to access this store",
        })
        return
    }

    result, err := service.GetAllProductStockLogs(storeUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Product stock logs fetched successfully",
        "data":    result,
    })
}

func GetProductStockLogsByProductHandler(c *gin.Context) {
    storeUUID := c.Param("store_uuid")
    productUUID := c.Param("product_uuid")

    if storeUUID == "" || productUUID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "store_uuid and product_uuid are required",
        })
        return
    }

    userUUID, exists := c.Get("user_uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Unauthorized",
        })
        return
    }

    owned, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
    if err != nil || !owned {
        c.JSON(http.StatusForbidden, gin.H{
            "status":  "error",
            "message": "You are not allowed to access this store",
        })
        return
    }

    logs, err := service.GetProductStockLogsByProduct(storeUUID, productUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to fetch product stock logs",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Product stock logs fetched successfully",
        "data":    logs,
    })
}