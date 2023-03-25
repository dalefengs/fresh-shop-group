package shop

import (
	"fresh-shop/server/global"
)

// GoodsImage 结构体
type GoodsImage struct {
	global.DbModel
	GoodsId *int   `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	Name    string `json:"name" form:"name" gorm:"column:name;default:'';comment:文件名"`
	Type    *int   `json:"type" form:"type" gorm:"column:type;default:0;comment:类型（0图片 1视频);"`
	Url     string `json:"url" form:"url" gorm:"column:url;comment:地址;size:500;"`
	Sort    *int   `json:"sort" form:"sort" gorm:"column:sort;default:50;comment:排序;size:10;"`
}

// TableName GoodsImage 表名
func (GoodsImage) TableName() string {
	return "shop_goods_image"
}
