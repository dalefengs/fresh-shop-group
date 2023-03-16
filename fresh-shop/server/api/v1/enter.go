package v1

import (
	"fresh-shop/server/api/v1/account"
	"fresh-shop/server/api/v1/business"
	"fresh-shop/server/api/v1/file"
	"fresh-shop/server/api/v1/shop"
	"fresh-shop/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  file.ApiGroup
	AccountApiGroup  account.ApiGroup
	BusinessApiGroup business.ApiGroup
	ShopApiGroup     shop.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
