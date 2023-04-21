package initialize

import (
	"fresh-shop/server/global"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
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
}
