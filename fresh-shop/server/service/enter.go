package service

import (
	"fresh-shop/server/service/account"
	"fresh-shop/server/service/business"
	"fresh-shop/server/service/file"
	"fresh-shop/server/service/shop"
	"fresh-shop/server/service/system"
	"fresh-shop/server/service/wechat"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  file.ServiceGroup
	AccountServiceGroup  account.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	ShopServiceGroup     shop.ServiceGroup
	WechatServiceGroup   wechat.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
