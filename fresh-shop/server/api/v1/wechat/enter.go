package wechat

import "fresh-shop/server/service"

type ApiGroup struct {
	WeChatApi
}

var (
	wechatService = service.ServiceGroupApp.WechatServiceGroup.WechatService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
)
