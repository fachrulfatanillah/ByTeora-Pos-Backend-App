package response

type CreateProductResponse struct {
	ProductUUID   string   `json:"product_uuid"`
	StoreUUID     string   `json:"store_uuid"`
	CategoryUUID  *string  `json:"category_uuid,omitempty"`
	ProductName   string   `json:"product_name"`
	SKU           *string  `json:"sku,omitempty"`
	Barcode       *string  `json:"barcode,omitempty"`
	Description   *string  `json:"description,omitempty"`
	Price         float64  `json:"price"`
	Cost          *float64 `json:"cost,omitempty"`
	Status        string   `json:"status"`
}

type ProductItemResponse struct {
	ProductUUID  string   `json:"product_uuid"`
	StoreUUID    string   `json:"store_uuid"`
	CategoryUUID *string  `json:"category_uuid,omitempty"`
	ProductName  string   `json:"product_name"`
	SKU          *string  `json:"sku,omitempty"`
	Barcode      *string  `json:"barcode,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Price        float64  `json:"price"`
	Cost         *float64 `json:"cost,omitempty"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type UpdateProductResponse struct {
	ProductUUID  string   `json:"product_uuid"`
	StoreUUID    string   `json:"store_uuid"`
	CategoryUUID *string  `json:"category_uuid,omitempty"`
	ProductName  string   `json:"product_name"`
	SKU          *string  `json:"sku,omitempty"`
	Barcode      *string  `json:"barcode,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Price        float64  `json:"price"`
	Cost         *float64 `json:"cost,omitempty"`
	Status       string   `json:"status"`
	UpdatedAt    string   `json:"updated_at"`
}