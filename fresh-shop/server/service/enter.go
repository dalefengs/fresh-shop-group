package service

import (
	"fresh-shop/server/service/account"
	"fresh-shop/server/service/example"
	"fresh-shop/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AccountServiceGroup account.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
