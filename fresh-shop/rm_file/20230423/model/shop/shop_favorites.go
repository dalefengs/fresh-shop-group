package shop

import (
	"fresh-shop/server/global"
)

// ShopFavorites 结构体
type ShopFavorites struct {
	global.DbModel
	GoodsId *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
}

// TableName ShopFavorites 表名
func (ShopFavorites) TableName() string {
	return "shop_favorites"
}
