package initialize

import (
	"fresh-shop/server/global"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
)

func Wechat() {
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     global.Config.Wechat.Appid,
		AppSecret: global.Config.Wechat.Secret,
		Cache:     memory,
	}
	wx := wechat.NewWechat()
	global.MiniProgram = wx.GetMiniProgram(cfg)

	// 初始化微信支付
	wxPayCfg := &payConfig.Config{
		AppID:     global.Config.Wechat.Appid,
		MchID:     global.Config.WechatPay.MchId,
		Key:       global.Config.WechatPay.ApiV2Key,
		NotifyURL: global.Config.WechatPay.NotifyURL,
	}
	wxPay := pay.NewPay(wxPayCfg)
	global.WxPay = wxPay
}
