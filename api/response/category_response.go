package response

type CreateCategoryResponse struct {
	CategoryUUID string `json:"category_uuid"`
	StoreUUID    string `json:"store_uuid"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}

type CategoryResponse struct {
	CategoryUUID string `json:"category_uuid"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}