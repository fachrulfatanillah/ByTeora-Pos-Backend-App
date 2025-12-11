package controller

import (
    "ByTeora-Pos-Backend-App/api/request"
    "ByTeora-Pos-Backend-App/api/response"
    "ByTeora-Pos-Backend-App/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

func CreateStore(c *gin.Context) {
    userUUID, exists := c.Get("user_uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "failed",
            "message": "Unauthorized access",
        })
        return
    }

    var req request.CreateStoreRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "failed",
            "message": "Invalid request body",
            "error":   err.Error(),
        })
        return
    }

    userID, err := repository.GetUserIDByUUID(userUUID.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "failed",
            "message": "User not found",
        })
        return
    }

    storeUUID := uuid.NewString()

    err = repository.CreateStore(
        storeUUID,
        userID,
        req.StoreName,
        req.Address,
        req.PhoneNumber,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "failed",
            "message": "Failed to create store",
        })
        return
    }
    res := response.CreateStoreResponse{
        StoreUUID:   storeUUID,
        UserUUID:    userUUID.(string),
        StoreName:   req.StoreName,
        Address:     req.Address,
        PhoneNumber: req.PhoneNumber,
        Status:      "active",
    }

    c.JSON(http.StatusCreated, gin.H{
        "status":  "success",
        "message": "Store created successfully",
        "data":    res,
    })
}

func GetStoresByUser(c *gin.Context) {
    var req request.GetStoreRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "failed",
            "message": "invalid request body",
        })
        return
    }

    stores, err := repository.GetStoresByUserUUID(req.UserUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "failed",
            "message": "Failed to fetch stores",
        })
        return
    }

    // Convert ke response struct
    var storeResponses []response.StoreResponse
    for _, s := range stores {
        storeResponses = append(storeResponses, response.StoreResponse{
            StoreUUID:   s.UUID,
            StoreName:   s.StoreName,
            Address:     s.Address,
            PhoneNumber: s.PhoneNumber,
            Status:      s.Status,
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Stores fetched successfully",
        "data":    storeResponses,
    })
}

func UpdateStore(c *gin.Context) {
    storeUUID := c.Param("store_uuid")

    userUUID, exists := c.Get("user_uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "failed",
            "message": "Unauthorized",
        })
        return
    }

    oldStore, err := repository.GetStoreByUUID(storeUUID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  "failed",
            "message": "Store not found",
        })
        return
    }

    belongs, err := repository.IsStoreOwnedByUser(storeUUID, userUUID.(string))
    if err != nil || !belongs {
        c.JSON(http.StatusForbidden, gin.H{
            "status": "failed",
            "message": "You cannot edit this store",
        })
        return
    }

    var req request.UpdateStoreRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "failed",
            "message": "Invalid request body",
            "error":   err.Error(),
        })
        return
    }

    if req.StoreName == "" {
        req.StoreName = oldStore.StoreName
    }
    if req.Address == "" {
        req.Address = oldStore.Address
    }
    if req.PhoneNumber == "" {
        req.PhoneNumber = oldStore.PhoneNumber
    }
    if req.Status == "" {
        req.Status = oldStore.Status
    }

    err = repository.UpdateStore(storeUUID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "failed",
            "message": "Failed to update store",
        })
        return
    }

    updatedStore, err := repository.GetStoreByUUID(storeUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "failed",
            "message": "Failed to fetch updated store",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Store updated successfully",
        "data": gin.H{
            "store_uuid":   updatedStore.UUID,
            "store_name":   updatedStore.StoreName,
            "address":      updatedStore.Address,
            "phone_number": updatedStore.PhoneNumber,
            "status":       updatedStore.Status,
        },
    })
}