package shop

import (
	"fresh-shop/server/global"
	"time"
)

// OrderReturn 结构体
type OrderReturn struct {
      global.DbModel
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单Id;size:20;"`
      Reason  string `json:"reason" form:"reason" gorm:"column:reason;comment:申请原因;size:255;"`
      Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:退款金额;size:14;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:售后状态(-1 拒绝售后 0未处理 1已退款);"`
      RefundStatus  *int `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款状态;"`
      Reply  string `json:"reply" form:"reply" gorm:"column:reply;comment:售后说明;size:255;"`
      ProcessTime  *time.Time `json:"processTime" form:"processTime" gorm:"column:process_time;comment:售后处理时间;"`
}


// TableName OrderReturn 表名
func (OrderReturn) TableName() string {
  return "shop_order_return"
}

