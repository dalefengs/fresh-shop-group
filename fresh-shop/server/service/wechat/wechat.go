package wechat

import (
	"errors"
	"fmt"
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/wechat/request"
	"fresh-shop/server/utils"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/silenceper/wechat/v2/pay/notify"
	orderPay "github.com/silenceper/wechat/v2/pay/order"
	"gorm.io/gorm"
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

type Scent struct {
	Uid        uint
	CreateTime string
}

// GetUnlimitedQRCode 获取小程序码 不限制
func (s *WechatService) GetUnlimitedQRCode(uid uint) (response []byte, err error) {
	code := global.MiniProgram.GetQRCode()
	param := qrcode.QRCoder{
		Scene:      fmt.Sprintf("uid=%d", uid),
		EnvVersion: "release",
	}
	response, err = code.GetWXACodeUnlimit(param)
	if err != nil {
		global.SugarLog.Errorf("获取小程序码失败, err:%s \n", err.Error())
		return
	}
	return
}

func JSAPIPay(openId, orderSn string, orderId uint, amount float64, createIp string) (err error, result *orderPay.Config) {
	order := global.WxPay.GetOrder()
	param := &orderPay.Params{
		OpenID:     openId,
		Body:       fmt.Sprintf("用户下单 金额:%.2f", amount),
		OutTradeNo: orderSn,
		TotalFee:   fmt.Sprintf("%.0f", amount*100), // 订单总金额，单位为分，详见支付金额
		CreateIP:   createIp,
		TimeExpire: time.Now().Add(30 * time.Minute).Format("20060102150405"), // 30 分钟过期
		TradeType:  "JSAPI",                                                   // 交易类型
		Attach:     strconv.Itoa(int(orderId)),                                // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	}
	if global.Config.WechatPay.Debug {
		param.Body = "测试支付"
		param.TotalFee = "1"
	}
	preOrder, err := order.BridgeConfig(param)
	if err != nil {
		global.SugarLog.Errorf("微信支付 - 发起 JSAPI 发生错误:%s", err.Error())
		return
	}
	result = &preOrder
	return
}

// NotifyLogic 支付回调逻辑处理
func (s *WechatService) NotifyLogic(req *notify.PaidResult) error {
	orderSn := *req.OutTradeNo
	log := fmt.Sprintf("订单支付回调逻辑: 订单号：%s, ", orderSn)
	var order shop.Order
	if errors.Is(global.DB.Where("order_sn = ?", orderSn).First(&order).Error, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf(log + "订单不存在 \n")
		return errors.New("订单不存在")
	}
	// 如果订单已经支付则直接结束
	if *order.Status == 1 {
		global.SugarLog.Errorf(log + "订单已支付 \n")
		return nil
	}
	if *order.Status != 0 {
		global.SugarLog.Errorf(log+"订单状态不正确, Status：%d \n", *order.Status)
		return errors.New("订单状态不正确")
	}
	finishStr := fmt.Sprintf("%.2f", float64(*req.TotalFee)/100)
	finish, err := strconv.ParseFloat(finishStr, 64)
	if err != nil {
		global.SugarLog.Errorf(log+"finishStr 转换 float64 失败, req.TotalFee:%s finishStr:%d \n", *req.TotalFee, finishStr)
		return err
	}
	// req.TimeEnd 转换为 time.Time 类型
	timeEnd, err := time.Parse("20060102150405", *req.TimeEnd)
	if err != nil {
		global.SugarLog.Errorf(log+"timeEnd 格式化时间失败, req.TimeEnd:%s \n", *req.TimeEnd)
		return err
	}
	order.Finish = finish
	order.Status = utils.Pointer(1)
	order.PayTime = &timeEnd
	order.PaymentOpenid = *req.OpenID
	order.PaymentInfo = *req.Attach
	order.TransationId = *req.TransactionID
	if err := global.DB.Save(&order).Error; err != nil {
		global.SugarLog.Errorf(log+"保存订单信息失败, err:%s \n", err.Error())
		return err
	}

	// 生成流水记录

	global.SugarLog.Infof(log + "支付成功")
	return nil
}
