package account

import (
	"fresh-shop/server/global"
	sysModel "fresh-shop/server/model/system"
)

// UserFinance 结构体
type UserFinance struct {
	global.DbModel
	UserId      *int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	Username    string           `json:"username" form:"username" gorm:"column:username;comment:用户名;size:191;"`
	TypeId      *int             `json:"typeId" form:"typeId" gorm:"column:type_id;comment:流水类型id;"`
	OptionType  *int             `json:"optionType" form:"optionType" gorm:"column:option_type;comment:操作类型(0余额 1冻结 2锁仓);"`
	IsFee       *int             `json:"isFee" form:"isFee" gorm:"column:is_fee;comment:是否手续费(0否 1是);"`
	FeeAmount   *float64         `json:"feeAmount" form:"feeAmount" gorm:"column:fee_amount;comment:手续费;size:14;"`
	Amount      *float64         `json:"amount" form:"amount" gorm:"column:amount;comment:变动数额;size:14;"`
	Balance     *float64         `json:"balance" form:"balance" gorm:"column:balance;comment:变动后账户余额;size:14;"`
	FromId      string           `json:"fromId" form:"fromId" gorm:"column:from_id;comment:来源id（类似订单Id，用于后台追溯）;size:50;"`
	FromUserId  *int             `json:"fromUserId" form:"fromUserId" gorm:"column:from_user_id;comment:来源用户id;size:20;"`
	FromName    string           `json:"fromName" form:"fromName" gorm:"column:from_name;comment:来源用户名;size:191;"`
	Remarks     string           `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:255;"`
	User        sysModel.SysUser `json:"user"`
	FinanceType UserFinanceType  `json:"financeType" gorm:"foreignKey:type_id"`
}
