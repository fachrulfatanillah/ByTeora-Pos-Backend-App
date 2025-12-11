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