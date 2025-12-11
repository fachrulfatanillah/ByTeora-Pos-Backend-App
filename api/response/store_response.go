package response

type CreateStoreResponse struct {
    StoreUUID   string `json:"store_uuid"`
    UserUUID    string `json:"user_uuid"`
    StoreName   string `json:"store_name"`
    Address     string `json:"address"`
    PhoneNumber string `json:"phone_number"`
    Status      string `json:"status"`
}

type StoreResponse struct {
    StoreUUID   string `json:"store_uuid"`
    StoreName   string `json:"store_name"`
    Address     string `json:"address"`
    PhoneNumber string `json:"phone_number"`
    Status      string `json:"status"`
}

type UpdateStoreResponse struct {
    StoreUUID   string `json:"store_uuid"`
    StoreName   string `json:"store_name"`
    Address     string `json:"address"`
    PhoneNumber string `json:"phone_number"`
    Status      string `json:"status"`
}