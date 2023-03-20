package shop

import (
	"fresh-shop/server/global"
)

// GoodsSpecItem 结构体
type GoodsSpecItem struct {
	global.DbModel
	SpecId *int   `json:"specId" form:"specId" gorm:"column:spec_id;comment:规格id;size:20;"`
	ImgUrl string `json:"imgUrl" form:"imgUrl" gorm:"column:img_url;comment:图片地址;size:500;"`
	Item   string `json:"item" form:"item" gorm:"column:item;comment:规格项;size:50;"`
	Sort   int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`

	ItemId string `gorm:"-"` // 用于记录创建规格时 前端的id，方便规格值查找
}

// TableName GoodsSpecItem 表名
func (GoodsSpecItem) TableName() string {
	return "shop_goods_spec_item"
}
