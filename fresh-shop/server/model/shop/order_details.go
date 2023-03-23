package shop

import (
	"fresh-shop/server/global"
	
)

// OrderDetails 结构体
type OrderDetails struct {
      global.DbModel
      GoodsId  *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
      GoodsName  string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:255;"`
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单Id;size:20;"`
      SpecId  *int `json:"specId" form:"specId" gorm:"column:spec_id;comment:规格id;size:20;"`
      SpecKeyName  string `json:"specKeyName" form:"specKeyName" gorm:"column:spec_key_name;comment:规格中文名(例：款式:香辣味,重量:200g);size:255;"`
      GoodsImage  string `json:"goodsImage" form:"goodsImage" gorm:"column:goods_image;comment:商品图片;size:255;"`
      Unit  string `json:"unit" form:"unit" gorm:"column:unit;comment:商品单位;size:10;"`
      Num  *int `json:"num" form:"num" gorm:"column:num;comment:商品数量;size:10;"`
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:订单价格;size:14;"`
      Total  *float64 `json:"total" form:"total" gorm:"column:total;comment:订单总价格;size:14;"`
      GiftPoints  *float64 `json:"giftPoints" form:"giftPoints" gorm:"column:gift_points;comment:赠送积分数量;size:10;"`
}


// TableName OrderDetails 表名
func (OrderDetails) TableName() string {
  return "shop_order_details"
}

