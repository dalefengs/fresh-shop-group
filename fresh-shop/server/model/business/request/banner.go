package request

import (
	"fresh-shop/server/model/business"
	"fresh-shop/server/model/common/request"
	"time"
)

type BannerSearch struct {
	business.Banner
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
