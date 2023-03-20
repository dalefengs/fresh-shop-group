package shop

import (
	"fresh-shop/server/global"
	
)

// GoodsDescription 结构体
type GoodsDescription struct {
      global.DbModel
      GoodsId  *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
      Notice  string `json:"notice" form:"notice" gorm:"column:notice;comment:购买须知;"`
      Details  string `json:"details" form:"details" gorm:"column:details;comment:商品详情;"`
}


// TableName GoodsDescription 表名
func (GoodsDescription) TableName() string {
  return "shop_goods_description"
}

