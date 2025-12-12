package request

type CreateProductRequest struct {
    CategoryUUID *string  `json:"category_uuid,omitempty"`
    ProductName  string   `json:"product_name" binding:"required"`
    SKU          *string  `json:"sku,omitempty"`
    Barcode      *string  `json:"barcode,omitempty"`
    Description  *string  `json:"description,omitempty"`
    Price        float64  `json:"price" binding:"required"`
    Cost         *float64 `json:"cost,omitempty"`
    Status       *string  `json:"status,omitempty"`
}

type UpdateProductRequest struct {
	CategoryUUID *string  `json:"category_uuid,omitempty"`
	ProductName  *string  `json:"product_name,omitempty"`
	SKU          *string  `json:"sku,omitempty"`
	Barcode      *string  `json:"barcode,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Price        *float64 `json:"price,omitempty"`
	Cost         *float64 `json:"cost,omitempty"`
	Status       *string  `json:"status,omitempty"`
}