package business

import (
	"fresh-shop/server/global"
	
)

// UserDelivery 结构体
type UserDelivery struct {
      global.DbModel
      Name  string `json:"name" form:"name" gorm:"column:name;comment:送货人姓名;size:255;"`
      Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:送货人手机号;size:11;"`
      DeliverCount  *int `json:"deliverCount" form:"deliverCount" gorm:"column:deliver_count;comment:送货单数;size:10;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态（0禁用 1启用）;"`
}


// TableName UserDelivery 表名
func (UserDelivery) TableName() string {
  return "user_delivery"
}

