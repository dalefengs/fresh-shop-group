package shop

import (
	"fresh-shop/server/global"
)

// GoodsSpec 结构体
type GoodsSpec struct {
	global.DbModel
	GoodsId       int    `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:20;"`
	Title         string `json:"title" form:"title" gorm:"column:title;comment:规格名称;size:50;"`
	IsUploadImage int    `json:"isUploadImage" form:"isUploadImage" gorm:"column:is_upload_image;comment:是否可上传图片（0否 1是）;"`
	Sort          int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`

	SpecId string `gorm:"-"` // 用于记录创建规格时 前端的id，方便规格值查找
}

// TableName GoodsSpec 表名
func (GoodsSpec) TableName() string {
	return "shop_goods_spec"
}
