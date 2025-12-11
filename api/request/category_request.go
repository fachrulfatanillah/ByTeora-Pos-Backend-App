package request

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
	Description  string `json:"description,omitempty"`
}

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name,omitempty"`
	Description  string `json:"description,omitempty"`
	Status       string `json:"status,omitempty"`
}