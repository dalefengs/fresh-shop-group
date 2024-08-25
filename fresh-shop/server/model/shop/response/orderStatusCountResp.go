package response

type OrderStatusCountResponse struct {
	Unpaid    int `json:"unpaid"`    // 待付款
	Delivered int `json:"delivered"` // 待发货
	Shipped   int `json:"shipped"`   // 已发货
	Success   int `json:"success"`

	OrderMonthStatistics
}
