package router

import (
	"fresh-shop/server/router/account"
	"fresh-shop/server/router/example"
	"fresh-shop/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Account account.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
