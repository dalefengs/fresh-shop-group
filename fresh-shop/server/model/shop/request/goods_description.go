package request

import (
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
	"time"
)

type GoodsDescriptionSearch struct{
    shop.GoodsDescription
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
