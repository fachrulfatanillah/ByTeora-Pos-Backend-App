package response

type ProductStockResponse struct {
	StockUUID    string `json:"stock_uuid"`
	ProductUUID  string `json:"product_uuid"`
	StoreUUID    string `json:"store_uuid"`
	StockIn      int    `json:"stock_in"`
	StockOut     int    `json:"stock_out"`
	CurrentStock int    `json:"current_stock"`
	Status       string `json:"status"`
}

type ProductStockLogResponse struct {
    StockUUID     string `json:"stock_uuid"`
    ProductUUID   string `json:"product_uuid"`
    StoreUUID     string `json:"store_uuid"`
    StockIn       int    `json:"stock_in"`
    StockOut      int    `json:"stock_out"`
    CurrentStock  int    `json:"current_stock"`
    Status        string `json:"status"`
    CreatedAt     string `json:"created_at"`
}