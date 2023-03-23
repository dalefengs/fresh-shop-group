package shop

import (
	"fresh-shop/server/global"
	"time"
)

// Order 结构体
type Order struct {
	global.DbModel
	UserId         *int           `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	GoodsId        *int           `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	OrderSn        string         `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:订单编号;size:50;"`
	GoodsArea      *int           `json:"goodsArea" form:"goodsArea" gorm:"column:goods_area;comment:所属区域(0普通 1积分商城);"`
	ShipmentName   string         `json:"shipmentName" form:"shipmentName" gorm:"column:shipment_name;comment:收货人姓名;size:20;"`
	ShipmentMobile string         `json:"shipmentMobile" form:"shipmentMobile" gorm:"column:shipment_mobile;comment:收货人手机号;size:11;"`
	ShipmentArea   string         `json:"shipmentArea" form:"shipmentArea" gorm:"column:shipment_area;comment:收货人地区编码;size:10;"`
	ShipingAddress string         `json:"shipingAddress" form:"shipingAddress" gorm:"column:shiping_address;comment:收获人地址;size:255;"`
	Num            *int           `json:"num" form:"num" gorm:"column:num;comment:商品总数量;size:10;"`
	Total          *float64       `json:"total" form:"total" gorm:"column:total;comment:订单商品总金额;size:14;"`
	Postage        *float64       `json:"postage" form:"postage" gorm:"column:postage;comment:邮费;size:14;"`
	Finish         *float64       `json:"finish" form:"finish" gorm:"column:finish;comment:实付金额;size:14;"`
	Payment        *int           `json:"payment" form:"payment" gorm:"column:payment;comment:支付方式(1余额 2微信);"`
	PaymentInfo    string         `json:"paymentInfo" form:"paymentInfo" gorm:"column:payment_info;comment:支付详情信息;size:255;"`
	PaymentOpenid  string         `json:"paymentOpenid" form:"paymentOpenid" gorm:"column:payment_openid;comment:支付openId;size:255;"`
	TransationId   string         `json:"transationId" form:"transationId" gorm:"column:transation_id;comment:支付流水订单号;size:255;"`
	Remarks        string         `json:"remarks" form:"remarks" gorm:"column:remarks;comment:留言;size:255;"`
	Status         *int           `json:"status" form:"status" gorm:"column:status;comment:订单状态(0未付款 1已付款待发货 2 已发货 已收货);"`
	StatusCancel   *int           `json:"statusCancel" form:"statusCancel" gorm:"column:status_cancel;comment:取消状态(0未取消 1用户取消 2后台取消 3超时取消);"`
	StatusRefund   *int           `json:"statusRefund" form:"statusRefund" gorm:"column:status_refund;comment:退款状态(0未退款 1退款中 2已退款 3退款失败);"`
	PayTime        *time.Time     `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;"`
	ShipmentTime   *time.Time     `json:"shipmentTime" form:"shipmentTime" gorm:"column:shipment_time;comment:发货时间;"`
	ReceiveTime    *time.Time     `json:"receiveTime" form:"receiveTime" gorm:"column:receive_time;comment:收货时间;"`
	CancelTime     *time.Time     `json:"cancelTime" form:"cancelTime" gorm:"column:cancel_time;comment:取消时间;"`
	OrderDetails   []OrderDetails `json:"orderDetails"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "shop_order"
}
