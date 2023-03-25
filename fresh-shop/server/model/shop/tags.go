package shop

import (
	"fresh-shop/server/global"
)

// Tags 结构体
type Tags struct {
	global.DbModel
	Name string `json:"name" form:"name" gorm:"column:name;comment:标签名;size:20;"`
	Sort *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName Tags 表名
func (Tags) TableName() string {
	return "shop_tags"
}
