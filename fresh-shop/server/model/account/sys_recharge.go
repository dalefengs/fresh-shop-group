package account

import (
	"fresh-shop/server/global"
	sysModel "fresh-shop/server/model/system"
)

// SysRecharge 结构体
type SysRecharge struct {
	global.DbModel
	UserId    *int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	Username  string           `json:"username" form:"username" gorm:"column:username;comment:用户名;size:191;"`
	GroupId   *int             `json:"groupId" form:"groupId" gorm:"column:group_id;comment:币种id;size:20;"`
	AdminId   *int             `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:管理员id;size:20;"`
	AdminName string           `json:"adminName" form:"adminName" gorm:"column:admin_name;comment:管理员用户名;size:191;"`
	Remarks   string           `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:255;"`
	Amount    *float64         `json:"amount" form:"amount" gorm:"column:amount;comment:增减数额;size:14;"`
	Balance   *float64         `json:"balance" form:"balance" gorm:"column:balance;comment:变动后数额;size:14;"`
	User      sysModel.SysUser `json:"user"`
	Account   Account          `json:"account" gorm:"foreignKey:user_id; references:user_id"`
	Group     AccountGroup     `json:"group" gorm:"foreignKey:group_id"`
}

// TableName SysRecharge 表名
func (SysRecharge) TableName() string {
	return "sys_recharge"
}
