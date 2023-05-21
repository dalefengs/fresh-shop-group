package shop

import (
	"fresh-shop/server/global"
)

// OrderReturnDetails 结构体
type OrderReturnDetails struct {
	global.DbModel
	OrderDetailId *int `json:"orderDetailId" form:"orderDetailId" gorm:"column:order_detail_id;comment:订单商品Id;size:20;"`
	ReturnId      *int `json:"returnId" form:"returnId" gorm:"column:return_id;comment:售后订单Id;size:20;"`
	Num           *int `json:"num" form:"num" gorm:"column:num;comment:申请售后的数量;size:10;"`
}

// TableName OrderReturnDetails 表名
func (OrderReturnDetails) TableName() string {
	return "shop_order_return_details"
}
