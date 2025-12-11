package response

import "time"

type CreateCategoryResponse struct {
	CategoryUUID string `json:"category_uuid"`
	StoreUUID    string `json:"store_uuid"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}

type CategoryResponse struct {
	UUID         string    `json:"uuid"`
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}

type UpdateCategoryResponse struct {
    CategoryUUID string `json:"category_uuid"`
    StoreUUID    string `json:"store_uuid"`
    CategoryName string `json:"category_name"`
    Description  string `json:"description,omitempty"`
    Status       string `json:"status"`
}

type DeleteCategoryResponse struct {
	CategoryUUID string `json:"category_uuid"`
	StoreUUID    string `json:"store_uuid"`
}