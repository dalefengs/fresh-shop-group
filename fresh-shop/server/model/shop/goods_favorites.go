package shop

import (
	"fresh-shop/server/global"
)

// GoodsFavorites 结构体
type GoodsFavorites struct {
	global.DbModel
	UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	GoodsId *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
}

// TableName GoodsFavorites 表名
func (GoodsFavorites) TableName() string {
	return "shop_favorites"
}
