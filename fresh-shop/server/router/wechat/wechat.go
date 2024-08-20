package wechat

import (
	v1 "fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type WechatRouter struct {
}

// InitWechatRouter 初始化 WechatRouter 路由信息
func (s *WechatRouter) InitWechatRouter(Router *gin.RouterGroup) {
	weChatRouter := Router.Group("wechat").Use(middleware.OperationRecord())
	weChatRouterNotRecord := Router.Group("wechat")
	var weChatApi = v1.ApiGroupApp.WechatApiGroup.WeChatApi
	{
		weChatRouter.POST("createPayData", weChatApi.CreatePayData) // 创建预支付订单，返回小程序发起支付的参数
	}
	{
		weChatRouterNotRecord.GET("getUnlimitedQRCode", weChatApi.GetUnlimitedQRCode) // 获取小程序码
	}
}

// InitWechatPublicRouter 初始化公共的 WechatRouter 路由信息
func (s *WechatRouter) InitWechatPublicRouter(Router *gin.RouterGroup) {
	weChatRouterWithoutRecord := Router.Group("wechat")
	var weChatApi = v1.ApiGroupApp.WechatApiGroup.WeChatApi
	{
		weChatRouterWithoutRecord.GET("code2Session", weChatApi.Code2Session) // 换取 Session
		weChatRouterWithoutRecord.POST("pay/notify", weChatApi.PayNotify)     // 支付成功回调
	}
}
