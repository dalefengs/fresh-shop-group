package shop

import (
	"fresh-shop/server/global"
)

// Favorites 结构体
type Favorites struct {
	global.DbModel
	GoodsId *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
}

// TableName Favorites 表名
func (Favorites) TableName() string {
	return "shop_favorites"
}
