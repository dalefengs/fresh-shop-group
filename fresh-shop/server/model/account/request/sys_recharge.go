package request

import (
	"fresh-shop/server/model/account"
	"fresh-shop/server/model/common/request"
	"time"
)

type SysRechargeSearch struct {
	account.SysRecharge
	request.UserSearch
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
