package account

import (
	"fresh-shop/server/global"
)

// SysRecharge 结构体
type SysRecharge struct {
	global.DbModel
	UserId    *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	Username  string   `json:"username" form:"username" gorm:"column:username;comment:用户名;size:191;"`
	CoinType  *int     `json:"coinType" form:"coinType" gorm:"column:coin_type;comment:币种id;size:20;"`
	AdminId   *int     `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:管理员id;size:20;"`
	AdminName string   `json:"adminName" form:"adminName" gorm:"column:admin_name;comment:管理员用户名;size:191;"`
	Remarks   string   `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:255;"`
	Amount    *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:增减数额;size:14;"`
}

// TableName SysRecharge 表名
func (SysRecharge) TableName() string {
	return "sys_recharge"
}
