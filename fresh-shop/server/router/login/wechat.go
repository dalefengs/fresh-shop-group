package login

import (
	v1 "fresh-shop/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WechatRouter struct {
}

// InitWechatRouter 初始化 WechatRouter 路由信息
func (s *WechatRouter) InitWechatRouter(Router *gin.RouterGroup) {
	//weChatRouter := Router.Group("wechat").Use(middleware.OperationRecord())
	//var weChatApi = v1.ApiGroupApp.LoginApiGroup.WeChatApi
	//{
	//}
}

// InitWechatPublicRouter 初始化公共的 WechatRouter 路由信息
func (s *WechatRouter) InitWechatPublicRouter(Router *gin.RouterGroup) {
	weChatRouterWithoutRecord := Router.Group("wechat")
	var weChatApi = v1.ApiGroupApp.LoginApiGroup.WeChatApi
	{
		weChatRouterWithoutRecord.GET("code2Session", weChatApi.Code2Session) // 根据ID获取Banner
	}
}
