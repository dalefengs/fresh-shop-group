package login

import "fresh-shop/server/service"

type ApiGroup struct {
	WeChatApi
}

var (
	wechatService = service.ServiceGroupApp.LoginServiceGroup.WechatService
)
