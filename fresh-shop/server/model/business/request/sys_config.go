package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
	"time"
)

type SysConfigSearch struct {
	system.SysConfig
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
