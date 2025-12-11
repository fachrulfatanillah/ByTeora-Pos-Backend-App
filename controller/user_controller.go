package controller

import (
    "ByTeora-Pos-Backend-App/api/request"
    "ByTeora-Pos-Backend-App/api/response"
    "ByTeora-Pos-Backend-App/service"
    "ByTeora-Pos-Backend-App/utils"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

func CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "invalid request body",
		})
		return
	}

	// Validation
	if req.Email == "" || req.Password == "" || req.NamaDepan == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "email, password, and nama_depan are required",
		})
		return
	}

	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "password must be at least 6 characters",
		})
		return
	}

	// Check email format
	if !utils.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "invalid email format",
		})
		return
	}

	// Check email exists
	exists, err := service.IsEmailExists(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  "error",
			Message: "failed checking email",
		})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "email already exists",
		})
		return
	}

	// Count users
	totalUsers, err := service.CountUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  "error",
			Message: "failed checking user count",
		})
		return
	}

	role := "owner"
	if totalUsers == 0 {
		role = "admin"
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  "error",
			Message: "failed to hash password",
		})
		return
	}

	userUUID := uuid.NewString()

	err = service.CreateUser(
		userUUID,
		req.Email,
		hashedPassword,
		req.NamaDepan,
		req.NamaBelakang,
		role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  "error",
			Message: "failed to create user",
		})
		return
	}

	// SUCCESS RESPONSE
	c.JSON(http.StatusCreated, response.BaseResponse{
		Status: "success",
		Data: response.UserResponse{
			UUID:  userUUID,
			Email: req.Email,
		},
	})
}

func AuthLogin(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "invalid request body",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Status:  "error",
			Message: "email and password are required",
		})
		return
	}

	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Status:  "error",
			Message: "invalid email or password",
		})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Status:  "error",
			Message: "invalid email or password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.UUID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  "error",
			Message: "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Status: "success",
		Data: response.UserResponse{
			UUID:  user.UUID,
			Email: user.Email,
			Token: token,
		},
	})
}