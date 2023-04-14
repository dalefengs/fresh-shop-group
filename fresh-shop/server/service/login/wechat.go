package login

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/login/request"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

type WechatService struct {
}

func (s *WechatService) Code2SessionKey(session request.Jscode2SessionReq) (auth.ResCode2Session, error) {
	wx := wechat.NewWechat()
	mini := wx.GetMiniProgram(global.MiniCfg)
	auth := mini.GetAuth()
	code2Session, err := auth.Code2Session(session.Jscode)
	if err != nil {
		return code2Session, err
	}
	return code2Session, nil
}
