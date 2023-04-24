package shop

import (
	"fresh-shop/server/global"
)

// Cart 结构体
type Cart struct {
	global.DbModel
	GoodsId    *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	UserId     *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	SpecType   *int `json:"specType" form:"specType" gorm:"column:spec_type;comment:商品规格(0单规格 1多规格);"`
	SpecItemId *int `json:"specItemId" form:"specItemId" gorm:"column:spec_item_id;comment:规格Id;size:20;"`
	Num        *int `json:"num" form:"num" gorm:"column:num;comment:商品数量;size:10;"`
}

// TableName Cart 表名
func (Cart) TableName() string {
	return "shop_cart"
}
