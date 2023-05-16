package v1

import (
	"fresh-shop/server/api/v1/account"
	"fresh-shop/server/api/v1/business"
	"fresh-shop/server/api/v1/file"
	"fresh-shop/server/api/v1/shop"
	"fresh-shop/server/api/v1/system"
	"fresh-shop/server/api/v1/wechat"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  file.ApiGroup
	AccountApiGroup  account.ApiGroup
	BusinessApiGroup business.ApiGroup
	ShopApiGroup     shop.ApiGroup
	WechatApiGroup   wechat.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
