// Package account 自动生成模板Account
package account

import (
	"fresh-shop/server/global"
)

// Account 结构体
type Account struct {
	global.DbModel
	UserId       *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`
	GroupId      *int     `json:"groupId" form:"groupId" gorm:"column:group_id;comment:币种id;size:19;"`
	Amount       *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:可用余额;size:14;"`
	FreezeAmount *float64 `json:"freezeAmount" form:"freezeAmount" gorm:"column:freeze_amount;comment:冻结余额;size:14;"`
	LockAmount   *float64 `json:"lockAmount" form:"lockAmount" gorm:"column:lock_amount;comment:锁仓金额;size:14;"`
	InAmount     *float64 `json:"inAmount" form:"inAmount" gorm:"column:in_amount;comment:累计入账数额;size:14;"`
	OutAmount    *float64 `json:"outAmount" form:"outAmount" gorm:"column:out_amount;comment:累计出账数额;size:14;"`
	Status       *int     `json:"status" form:"status" gorm:"column:status;comment:账户状态;"`
}

// TableName Account 表名
func (Account) TableName() string {
	return "user_account"
}
