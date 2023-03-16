package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	"time"
)

type BrandCategorySearch struct {
	shop.BrandCategory
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	BrandIds       []int      `json:"brandIds" form:"brandIds"`
	request.PageInfo
}
