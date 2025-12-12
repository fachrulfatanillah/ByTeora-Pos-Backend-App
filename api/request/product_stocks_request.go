package request

type CreateProductStockRequest struct {
	StockIn  *int   `json:"stock_in,omitempty"`
	StockOut *int   `json:"stock_out,omitempty"`
}