package shop

import (
	"fresh-shop/server/global"
)

// Category 结构体
type Category struct {
	global.DbModel
	Pid     *int    `json:"pid" form:"pid" gorm:"column:pid;comment:上级分类id;size:20;"`
	Title   string  `json:"title" form:"title" gorm:"column:title;comment:分类名;size:10;"`
	ImgUrl  string  `json:"imgUrl" form:"imgUrl" gorm:"column:img_url;comment:分类图片;size:500;"`
	Soft    *int    `json:"soft" form:"soft" gorm:"column:soft;comment:排序;size:10;"`
	IsFirst *int    `json:"isFirst" form:"isFirst" gorm:"column:is_first;comment:是否首页(0否 1是);"`
	Brands  []Brand `json:"brands" gorm:"many2many:shop_brand_category"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "shop_category"
}
