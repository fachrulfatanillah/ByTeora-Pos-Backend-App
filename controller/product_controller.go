package controller

import (
	"net/http"

	"ByTeora-Pos-Backend-App/api/request"
	"ByTeora-Pos-Backend-App/api/response"
	"ByTeora-Pos-Backend-App/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	storeUUID := c.Param("store_uuid")
	if storeUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "store_uuid is required",
		})
		return
	}

	userUUID, exists := c.Get("user_uuid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized",
		})
		return
	}

	belongs, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
	if err != nil || !belongs {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "failed",
			"message": "You are not allowed to add product to this store",
		})
		return
	}

	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	storeID, err := service.GetStoreIDByUUID(storeUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Store not found",
		})
		return
	}

	var categoryID *int
	if req.CategoryUUID != nil {
		id, err := service.GetCategoryIDByUUID(*req.CategoryUUID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "Category not found",
			})
			return
		}
		categoryID = &id
	}

	productUUID := uuid.NewString()

	status := "active"
	if req.Status != nil {
		status = *req.Status
	}

	err = service.CreateProduct(productUUID, storeID, categoryID, req, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to create product",
		})
		return
	}

	var categoryUUID *string
	if req.CategoryUUID != nil {
		categoryUUID = req.CategoryUUID
	}

	res := response.CreateProductResponse{
		ProductUUID:  productUUID,
		StoreUUID:    storeUUID,
		CategoryUUID: categoryUUID,
		ProductName:  req.ProductName,
		SKU:          req.SKU,
		Barcode:      req.Barcode,
		Description:  req.Description,
		Price:        req.Price,
		Cost:         req.Cost,
		Status:       status,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Product created successfully",
		"data":    res,
	})
}

func GetAllProducts(c *gin.Context) {
	storeUUID := c.Param("store_uuid")
	if storeUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "store_uuid is required",
		})
		return
	}

	userUUID, exists := c.Get("user_uuid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized",
		})
		return
	}

	belongs, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
	if err != nil || !belongs {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "failed",
			"message": "You are not allowed to access this store's products",
		})
		return
	}

	products, err := service.GetProductsByStoreUUID(storeUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Products fetched successfully",
		"data": gin.H{
			"products": products,
		},
	})
}

func UpdateProductHandler(c *gin.Context) {
	storeUUID := c.Param("store_uuid")
	productUUID := c.Param("product_uuid")

	if storeUUID == "" || productUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "store_uuid and product_uuid are required",
		})
		return
	}

	userUUID, exists := c.Get("user_uuid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized",
		})
		return
	}

	belongs, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
	if err != nil || !belongs {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "failed",
			"message": "You are not allowed to access this store",
		})
		return
	}

	storeID, err := service.GetStoreIDByUUID(storeUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Store not found",
		})
		return
	}

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
		return
	}

	_, err = service.GetProductByUUID(productUUID, storeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Product not found",
		})
		return
	}

	err = service.UpdateProductPartial(productUUID, storeID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	updated, _ := service.GetProductByUUID(productUUID, storeID)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product updated successfully",
		"data":    updated,
	})
}

func DeleteProductHandler(c *gin.Context) {
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

	belongs, err := service.IsStoreOwnedByUser(storeUUID, userUUID.(string))
	if err != nil || !belongs {
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

	existsProduct, err := service.IsProductBelongsToStore(productUUID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to check product",
		})
		return
	}

	if !existsProduct {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Product not found",
		})
		return
	}

	err = service.SoftDeleteProduct(productUUID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}