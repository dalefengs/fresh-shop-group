package wechat

import (
	"fmt"
	"fresh-shop/server/global"
	"fresh-shop/server/model/wechat/request"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	orderPay "github.com/silenceper/wechat/v2/pay/order"
	"strconv"
	"time"
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

func (s *WechatService) CreatePayData(payReq request.WechatPayReq) error {
	return nil
}

func JSAPIPay(opeid, orderSn string, orderId uint, amount float64, createIp string) (err error, result *orderPay.PreOrder) {
	order := global.WxPay.GetOrder()
	param := &orderPay.Params{
		OpenID:     opeid,
		Body:       fmt.Sprintf("用户下单 金额:%.2f", amount),
		OutTradeNo: orderSn,
		TotalFee:   fmt.Sprintf("%.0f", amount*1000), // 订单总金额，单位为分，详见支付金额
		CreateIP:   createIp,
		TimeExpire: time.Now().Add(30 * time.Minute).Format("20060102150405"), // 30 分钟过期
		TradeType:  "JSAPI",                                                   // 交易类型
		Attach:     strconv.Itoa(int(orderId)),                                // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	}
	preOrder, err := order.PrePayOrder(param)
	if err != nil {
		global.SugarLog.Errorf("微信支付 - 发起 JSAPI 发生错误:%s", err.Error())
		return
	}
	result = &preOrder
	return
}
