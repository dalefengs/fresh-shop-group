package shop

import (
	"fresh-shop/server/global"
)

// Brand 结构体
type Brand struct {
	global.DbModel
	Name string `json:"name" form:"name" gorm:"column:name;comment:品牌名称;size:20;"`
	Logo string `json:"logo" form:"logo" gorm:"column:logo;comment:品牌LOGO;size:500;"`
	Sort *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName Brand 表名
func (Brand) TableName() string {
	return "shop_brand"
}
