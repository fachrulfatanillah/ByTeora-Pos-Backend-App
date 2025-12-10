package controller

import (
    "ByTeora-Pos-Backend-App/dto"
    "ByTeora-Pos-Backend-App/repository"
    "ByTeora-Pos-Backend-App/utils"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

func CreateUser(c *gin.Context) {
    var req dto.CreateUserRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Validation
    if req.Email == "" || req.Password == "" || req.NamaDepan == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email, password, and nama_depan are required"})
        return
    }

    if len(req.Password) < 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password is too short, minimum 6 characters"})
        return
    }

    // Check email format
    if !utils.IsValidEmail(req.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
        return
    }

    // Check if email already exists
    exists, err := repository.IsEmailExists(req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed checking email"})
        return
    }
    if exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
        return
    }

    // Count users → menentukan role pertama
    totalUsers, err := repository.CountUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed checking user count"})
        return
    }

    role := "owner"
    if totalUsers == 0 {
        role = "admin"
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
        return
    }

    // Generate UUID
    userUUID := uuid.NewString()

    // Insert to DB
    err = repository.CreateUser(
        userUUID,
        req.Email,
        hashedPassword,
        req.NamaDepan,
        req.NamaBelakang,
        role,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"uuid":  userUUID,
			"email": req.Email,
		},
	})
}