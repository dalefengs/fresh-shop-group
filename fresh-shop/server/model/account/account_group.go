// 自动生成模板AccountGroup
package account

import (
	"fresh-shop/server/global"
)

// AccountGroup 结构体
type AccountGroup struct {
	global.GVA_MODEL
	NameEn string   `json:"nameEn" form:"nameEn" gorm:"column:name_en;comment:币中英文名;size:20;"`
	NameCn string   `json:"nameCn" form:"nameCn" gorm:"column:name_cn;comment:币种中文名;size:20;"`
	Places *float64 `json:"places" form:"places" gorm:"column:places;comment:小数点位数;"`
	Status *int     `json:"status" form:"status" gorm:"column:status;comment:状态(0禁用 1启用);"`
	Sync   *int     `json:"sync" form:"sync" gorm:"column:sync;comment:同步状态(0未同步 1已同步) 生成用户对应的币种账户和币种流水表;"`
}

// TableName AccountGroup 表名
func (AccountGroup) TableName() string {
	return "user_account_group"
}
