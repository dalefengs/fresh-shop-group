package request

import (
	"fresh-shop/server/model/business"
	"fresh-shop/server/model/common/request"
	"time"
)

type UserDeliverySearch struct{
    business.UserDelivery
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
