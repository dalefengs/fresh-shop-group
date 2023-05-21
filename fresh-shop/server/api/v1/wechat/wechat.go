package wechat

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/wechat/request"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/pay/notify"
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

// CreatePayData 创建微信支付数据  暂时无用
func (w *WeChatApi) CreatePayData(c *gin.Context) {
	var data request.WechatPayReq
	err := c.ShouldBindQuery(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.OpenId == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	data.ClientIP = c.ClientIP()
	err = wechatService.CreatePayData(data)
	if err != nil {
		global.Log.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		//response.OkWithData(result, c)
	}
}

// PayNotify 支付完成回调
func (w *WeChatApi) PayNotify(c *gin.Context) {
	global.SugarLog.Infof("微信支付回调 开始 \n")
	var req notify.PaidResult
	err := c.ShouldBindXML(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	n := global.WxPay.GetNotify()
	if ok := n.PaidVerifySign(req); !ok {
		global.SugarLog.Errorf("支付回调签名验证失败! %#v \n", req)
		response.WxpayNotify("FAIL", "签名失败", c)
		return
	}

	if *req.ReturnCode != "SUCCESS" {
		global.SugarLog.Errorf("支付回调失败! ReturnCode:%s, ReturnMsg: %s \n", *req.ReturnCode, *req.ReturnMsg)
		response.WxpayNotify("FAIL", "参数格式校验错误", c)
		return
	}
	if *req.ResultCode != "SUCCESS" {
		global.SugarLog.Errorf("支付回调失败! ResultCode:%s, ResultMsg: %s \n", *req.ResultCode, *req.ErrCodeDes)
		response.WxpayNotify("FAIL", "ResultMsg: "+*req.ReturnMsg, c)
		return
	}
	global.SugarLog.Infof("微信支付回调 验证通过开始执行业务逻辑\n")
	err = wechatService.NotifyLogic(&req)
	if err != nil {
		global.SugarLog.Errorf("支付回调失败! err: %v \n", err)
		response.WxpayNotify("FAIL", "参数格式校验错误", c)
	} else {
		response.WxpayNotify("SUCCESS", "OK", c)
	}
}
