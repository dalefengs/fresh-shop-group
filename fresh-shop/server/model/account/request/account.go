package request

import (
	"fresh-shop/server/model/account"
	"fresh-shop/server/model/common/request"
	"time"
)

type AccountSearch struct {
	account.Account
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
