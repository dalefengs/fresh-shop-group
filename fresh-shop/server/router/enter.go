package router

import (
	"fresh-shop/server/router/account"
	"fresh-shop/server/router/business"
	"fresh-shop/server/router/file"
	"fresh-shop/server/router/shop"
	"fresh-shop/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  file.RouterGroup
	Account  account.RouterGroup
	Shop     shop.RouterGroup
	Business business.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
