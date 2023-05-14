package shop

import (
	"fresh-shop/server/global"
)

// UserAddress 结构体
type UserAddress struct {
	global.DbModel
	UserId    *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	IsDefault *int     `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否默认;"`
	Name      string   `json:"name" form:"name" gorm:"column:name;comment:收货人姓名;size:20;"`
	Mobile    string   `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:11;"`
	Area      string   `json:"area" form:"area" gorm:"column:area;comment:收货人地区编码;size:20;"`
	Address   string   `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`
	Title     string   `json:"title" form:"title" gorm:"column:title;comment:标志点位置名称;size:255;"`
	Detail    string   `json:"detail" form:"detail" gorm:"column:detail;comment:详细地址;size:255;"`
	Lable     string   `json:"lable" form:"lable" gorm:"column:lable;default:'';comment:标签;size:20;"`
	Sex       *int     `json:"sex" form:"sex" gorm:"column:sex;default:1;comment:性别;size:20;"`
	Longitude *float64 `json:"longitude" form:"longitude" gorm:"column:longitude;default:0;comment:经度;size:20;"`
	Latitude  *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;default:0;comment:纬度;size:20;"`
}

// TableName UserAddress 表名
func (UserAddress) TableName() string {
	return "user_address"
}
