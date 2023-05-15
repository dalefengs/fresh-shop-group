package system

import (
	"fresh-shop/server/global"
)

// SysConfig 结构体
type SysConfig struct {
	global.DbModel
	Name      string `json:"name" form:"name" gorm:"column:name;comment:配置参数的 key;size:20;"`
	Value     string `json:"value" form:"value" gorm:"column:value;comment:配置参数的 value;size:100;"`
	GroupType string `json:"groupType" form:"groupType" gorm:"column:group_type;comment:配置分组;size:20;"`
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:中文描述;size:100;"`
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:状态(1开启 0禁用);"`
}

// TableName SysConfig 表名
func (SysConfig) TableName() string {
	return "sys_config"
}
