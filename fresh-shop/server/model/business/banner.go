package business

import (
	"fresh-shop/server/global"
)

// Banner 结构体
type Banner struct {
	global.DbModel
	ImgUrl string `json:"imgUrl" form:"imgUrl" gorm:"column:img_url;comment:图片Url;size:500;"`
	ToPath string `json:"toPath" form:"toPath" gorm:"column:to_path;comment:跳转地址;size:255;"`
	Type   *int   `json:"type" form:"type" gorm:"column:type;comment:跳转类型(0页面跳转);"`
	Sort   int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
}

// TableName Banner 表名
func (Banner) TableName() string {
	return "shop_banner"
}
