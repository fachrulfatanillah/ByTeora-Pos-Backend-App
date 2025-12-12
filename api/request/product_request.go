package request

type CreateProductRequest struct {
    CategoryID  *int     `json:"category_id,omitempty"`
    ProductName string   `json:"product_name" binding:"required"`
    SKU         *string  `json:"sku,omitempty"`
    Barcode     *string  `json:"barcode,omitempty"`
    Description *string  `json:"description,omitempty"`
    Price       float64  `json:"price" binding:"required"`
    Cost        *float64 `json:"cost,omitempty"`
    Status      *string  `json:"status,omitempty"`
}