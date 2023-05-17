package response

import (
	"fresh-shop/server/model/shop"
	orderPay "github.com/silenceper/wechat/v2/pay/order"
)

// CreateOrderResp 创建订单响应
type CreateOrderResp struct {
	Pay   orderPay.Config `json:"pay" form:"pay"`
	Order shop.Order      `json:"order" form:"order"`
}
