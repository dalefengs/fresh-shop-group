package login

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/login/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WeChatApi struct {
}

// Code2Session 换取用登陆凭证code调用后端接口去获取用户登录态信息，
// 包括用户在当前小程序的唯一标识（openid）、
// 微信开放平台帐号下的唯一标识（unionid，若当前小程序已绑定到微信开放平台帐号）
// 及本次登录的会话密钥（session_key）等（用户无感知）
func (w *WeChatApi) Code2Session(c *gin.Context) {
	var data request.Jscode2SessionReq
	err := c.ShouldBindQuery(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.Jscode == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	result, err := wechatService.Code2SessionKey(data)
	if err != nil {
		global.Log.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(result, c)
	}
}
