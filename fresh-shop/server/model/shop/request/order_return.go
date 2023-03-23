package request

import (
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
	"time"
)

type OrderReturnSearch struct{
    shop.OrderReturn
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartProcessTime  *time.Time  `json:"startProcessTime" form:"startProcessTime"`
    EndProcessTime  *time.Time  `json:"endProcessTime" form:"endProcessTime"`
    request.PageInfo
}
