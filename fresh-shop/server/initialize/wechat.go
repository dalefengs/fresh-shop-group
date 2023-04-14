package initialize

import (
	"fresh-shop/server/global"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func Wechat() {
	memory := cache.NewMemory()
	global.MiniCfg = &miniConfig.Config{
		AppID:     global.Config.Wechat.Appid,
		AppSecret: global.Config.Wechat.Secret,
		Cache:     memory,
	}
}
