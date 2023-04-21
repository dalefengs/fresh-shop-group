package login

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/login/request"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

type WechatService struct {
}

func (s *WechatService) Code2SessionKey(session request.Jscode2SessionReq) (auth.ResCode2Session, error) {
	a := global.MiniProgram.GetAuth()
	code2Session, err := a.Code2Session(session.Jscode)
	if err != nil {
		return code2Session, err
	}
	return code2Session, nil
}
