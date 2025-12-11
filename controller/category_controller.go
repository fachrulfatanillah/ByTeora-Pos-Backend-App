package controller

import (
	"ByTeora-Pos-Backend-App/api/request"
	"ByTeora-Pos-Backend-App/api/response"
	"ByTeora-Pos-Backend-App/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCategory(c *gin.Context) {
	storeUUID := c.Param("store_uuid")
	if storeUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Store UUID is required",
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed checking store ownership",
		})
		return
	}
	if !belongs {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "failed",
			"message": "You are not allowed to add category to this store",
		})
		return
	}

	var req request.CreateCategoryRequest
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Store not found",
		})
		return
	}

	categoryUUID := uuid.NewString()

	err = service.CreateCategory(categoryUUID, storeID, req.CategoryName, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to create category",
		})
		return
	}

	res := response.CreateCategoryResponse{
		CategoryUUID: categoryUUID,
		StoreUUID:    storeUUID,
		CategoryName: req.CategoryName,
		Description:  req.Description,
		Status:       "active",
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Category created successfully",
		"data":    res,
	})
}

func GetCategoriesByStore(c *gin.Context) {
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed checking store ownership",
		})
		return
	}

	if !belongs {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "failed",
			"message": "You are not allowed to get categories of this store",
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

	categories, err := service.GetCategoriesByStoreID(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to fetch categories",
		})
		return
	}

	var res []response.CategoryResponse
	for _, cat := range categories {
		res = append(res, response.CategoryResponse{
			UUID:         cat.UUID,
			CategoryName: cat.CategoryName,
			Description:  cat.Description,
			Status:       cat.Status,
			CreatedAt:    cat.CreatedAt,
			ModifiedAt:   cat.ModifiedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Categories fetched successfully",
		"data":    res,
	})
}