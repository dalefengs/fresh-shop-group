package initialize

import (
	_ "fresh-shop/server/source/example"
	_ "fresh-shop/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
