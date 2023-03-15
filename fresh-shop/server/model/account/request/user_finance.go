package request

import (
	"fresh-shop/server/model/account"
	"fresh-shop/server/model/common/request"
	"time"
)

type UserFinanceSearch struct {
	account.UserFinance
	request.UserSearch
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	GroupNameEn    string     `json:"groupNameEn" form:"groupNameEn"`
	request.PageInfo
}
