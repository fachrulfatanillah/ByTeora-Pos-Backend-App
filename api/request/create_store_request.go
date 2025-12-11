package request

type CreateStoreRequest struct {
    StoreName   string `json:"store_name" binding:"required"`
    Address     string `json:"address" binding:"required"`
    PhoneNumber string `json:"phone_number" binding:"required"`
}

type GetStoreRequest struct {
    UserUUID string `json:"user_uuid" binding:"required"`
}