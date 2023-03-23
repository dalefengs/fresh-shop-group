package request

import (
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
	"time"
)

type OrderDeliverySearch struct{
    shop.OrderDelivery
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartScheduledTime  *time.Time  `json:"startScheduledTime" form:"startScheduledTime"`
    EndScheduledTime  *time.Time  `json:"endScheduledTime" form:"endScheduledTime"`
    StartReceiptTime  *time.Time  `json:"startReceiptTime" form:"startReceiptTime"`
    EndReceiptTime  *time.Time  `json:"endReceiptTime" form:"endReceiptTime"`
    request.PageInfo
}
